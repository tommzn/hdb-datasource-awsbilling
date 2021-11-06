package awsbilling

import (
	"errors"
	"regexp"

	awsevents "github.com/aws/aws-lambda-go/events"
	"github.com/golang/protobuf/proto"
	core "github.com/tommzn/hdb-datasource-core"
)

// New creates a processor to generate billing reports.
func New(skipHeadline bool) core.S3EventProcessor {
	return &BillingReportProcessor{
		skipHeadline: skipHeadline,
	}
}

// Process is called to process given event for a S3 object.
// If download option is enable via config it will pass S3 object content as well.
func (processor *BillingReportProcessor) ProcessEvent(entity awsevents.S3Entity, content []byte) (proto.Message, error) {

	if !isBillingReportFile(entity.Object.Key) {
		return nil, errors.New("Not an AWS  billing report: " + entity.Object.Key)
	}

	lines, err := readObject(entity.Object.Key, content, processor.skipHeadline)
	if err != nil {
		return nil, err
	}

	billingItems, err := toBillingItems(lines)
	if err != nil {
		return nil, err
	}
	return generateBillingReport(billingItems)
}

func isBillingReportFile(s3Key string) bool {
	match, _ := regexp.MatchString(".*BillingReport.*csv.*", s3Key)
	return match
}
