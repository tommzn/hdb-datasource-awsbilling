package awsbilling

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConverterTestSuite struct {
	suite.Suite
}

func TestConverterTestSuite(t *testing.T) {
	suite.Run(t, new(ConverterTestSuite))
}

func (suite *ConverterTestSuite) TestCsvReportConversion() {

	objectKey := "fixtures/BillingReport.csv"
	content, err := readFixtures(objectKey)
	suite.Nil(err)
	lines, err := readObject(objectKey, content, true)
	suite.Nil(err)

	billingItems, err := toBillingItems(lines)
	suite.Nil(err)
	suite.Len(billingItems, 6)
}

func (suite *ConverterTestSuite) TestCsvReportNonCsv() {

	billingItems, err := toBillingItems([]string{""})
	suite.NotNil(err)
	suite.Len(billingItems, 0)
}

func (suite *ConverterTestSuite) TestCsvReportInvalidValues() {

	testData1 := ",,,,,,2018-11-01T00:00:00Z,,,,,,,,,,,,,,,,,Invalid Costs,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"
	testData2 := ",,,,,,Invalid DAte,,,,,,,,,,,,,,,,,Invalid Costs,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"
	billingItems, err := toBillingItems([]string{testData1, testData2})
	suite.NotNil(err)
	suite.Len(billingItems, 0)
}
