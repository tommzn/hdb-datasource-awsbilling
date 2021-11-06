package awsbilling

import "time"

// BillingReportProcessor reads context from a AWS billing report and generates a
// biling report events with sumarized amounts.
type BillingReportProcessor struct {
	skipHeadline bool
}

// billingItem contains values from a single billing report line.
type billingItem struct {
	billingPeriod time.Time
	itemType      string
	productCode   string
	currencyCode  string
	costs         float64
	region        string
}
