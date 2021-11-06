package awsbilling

import (
	"io/ioutil"
	"time"

	awsevents "github.com/aws/aws-lambda-go/events"
)

func readFixtures(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func billingItemsForTest() []billingItem {
	return []billingItem{
		billingItem{
			billingPeriod: time.Now(),
			itemType:      "Tax",
			productCode:   "Tax",
			currencyCode:  "USD",
			costs:         1.23,
			region:        "",
		},
		billingItem{
			billingPeriod: time.Now(),
			itemType:      "Credit",
			productCode:   "EC2",
			currencyCode:  "USD",
			costs:         3.45345,
			region:        "eu-central-1",
		},
		billingItem{
			billingPeriod: time.Now(),
			itemType:      "Credit",
			productCode:   "EC2",
			currencyCode:  "USD",
			costs:         3.42654,
			region:        "eu-central-1",
		},
	}
}

func s3EntityForTest(objectKey string) awsevents.S3Entity {
	return awsevents.S3Entity{
		Object: awsevents.S3Object{
			Key: objectKey,
		},
	}
}
