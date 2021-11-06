package awsbilling

import (
	"errors"
	"math"
	"strings"

	events "github.com/tommzn/hdb-events-go"
)

func generateBillingReport(billingItems []billingItem) (*events.BillingReport, error) {

	if len(billingItems) == 0 {
		return nil, errors.New("Empty list of billing items passed.")
	}

	report := &events.BillingReport{
		BillingPeriod: billingItems[0].billingPeriod.Format("Jan 2006"),
		BillingAmount: make(map[string]float64),
		TaxAmount:     make(map[string]float64),
	}
	for _, billingItem := range billingItems {

		if strings.ToUpper(billingItem.itemType) == "TAX" {
			addTaxAmount(report, billingItem.currencyCode, billingItem.costs)
		} else {
			addBillingAmount(report, billingItem.currencyCode, billingItem.costs)
		}
	}

	for currencyCode, amount := range report.BillingAmount {
		report.BillingAmount[currencyCode] = round(amount, .5, 2)
	}
	for currencyCode, amount := range report.TaxAmount {
		report.TaxAmount[currencyCode] = round(amount, .5, 2)
	}
	return report, nil
}

func addTaxAmount(report *events.BillingReport, currencyCode string, amount float64) {
	if _, ok := report.TaxAmount[currencyCode]; !ok {
		report.TaxAmount[currencyCode] = 0
	}
	report.TaxAmount[currencyCode] += amount
}

func addBillingAmount(report *events.BillingReport, currencyCode string, amount float64) {
	if _, ok := report.BillingAmount[currencyCode]; !ok {
		report.BillingAmount[currencyCode] = 0
	}
	report.BillingAmount[currencyCode] += amount
}

func round(val float64, roundOn float64, places int) float64 {

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
