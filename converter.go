package awsbilling

import (
	"encoding/csv"
	"strconv"
	"strings"
	"time"
)

/**

List of relevat columns in billing report
	6: 	bill/BillingPeriodStartDate
	9: 	lineItem/LineItemType
	12: lineItem/ProductCode
	19: lineItem/CurrencyCode
	21: lineItem/UnblendedCost
	23: lineItem/BlendedCost
	47: product/location
	62: product/region

*/

const COLUMN_BILLING_PERIOD = 7
const COLUMN_LINEITEMTYPE = 10
const COLUMN_PRODUCTCODE = 13
const COLUMN_CURRENCYCODE = 20
const COLUMN_COSTS = 24
const COLUMN_REGION = 63

// toBillingItems converts given content into billing items.
func toBillingItems(billingReportLines []string) ([]billingItem, error) {

	billingItems := []billingItem{}
	errors := []error{}
	for _, billingReportLine := range billingReportLines {

		reader := csv.NewReader(strings.NewReader(billingReportLine))
		columns, err := reader.Read()
		if err != nil {
			errors = append(errors, err)
		}
		if len(columns) < COLUMN_REGION {
			continue
		}

		period, err := time.Parse(time.RFC3339, columns[COLUMN_BILLING_PERIOD])
		if err != nil {
			errors = append(errors, err)
			continue
		}

		costs, err := strconv.ParseFloat(columns[COLUMN_COSTS], 64)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		billingItems = append(billingItems, billingItem{
			billingPeriod: period,
			itemType:      columns[COLUMN_LINEITEMTYPE],
			productCode:   columns[COLUMN_PRODUCTCODE],
			currencyCode:  columns[COLUMN_CURRENCYCODE],
			costs:         costs,
			region:        columns[COLUMN_REGION],
		})
	}
	return billingItems, asError(errors)
}
