package awsbilling

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BillingReportProcessorTestSuite struct {
	suite.Suite
}

func TestBillingReportProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(BillingReportProcessorTestSuite))
}

func (suite *BillingReportProcessorTestSuite) TestGenerateBillingReport() {

	processor := New(true)

	objectKey1 := "fixtures/BillingReport.csv"
	content1, err1 := readFixtures(objectKey1)
	suite.Nil(err1)
	report1, err1 := processor.ProcessEvent(s3EntityForTest(objectKey1), content1)
	suite.Nil(err1)
	suite.NotNil(report1)

	objectKey2 := "fixtures/BillingReport.csv.gz"
	content2, err2 := readFixtures(objectKey2)
	suite.Nil(err2)
	report2, err2 := processor.ProcessEvent(s3EntityForTest(objectKey2), content2)
	suite.Nil(err2)
	suite.NotNil(report2)

	emptyFile := "fixtures/empty.BillingReport.csv"
	content3, err3 := readFixtures(emptyFile)
	suite.Nil(err3)
	report3, err3 := processor.ProcessEvent(s3EntityForTest(emptyFile), content3)
	suite.Nil(err3)
	suite.NotNil(report3)

	testData1 := ",,,,,,2018-11-01T00:00:00Z,,,,,,,,,,,,,,,,,Invalid Costs,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"
	testData2 := ",,,,,,Invalid DAte,,,,,,,,,,,,,,,,,Invalid Costs,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"
	content4 := fmt.Sprintf("%s\n%s", testData1, testData2)
	report4, err4 := processor.ProcessEvent(s3EntityForTest("fixtures/BillingReport.csv"), []byte(content4))
	suite.NotNil(err4)
	suite.Nil(report4)

	report5, err5 := processor.ProcessEvent(s3EntityForTest(objectKey2), content1)
	suite.NotNil(err5)
	suite.Nil(report5)

	report6, err6 := processor.ProcessEvent(s3EntityForTest("invalid_file"), content1)
	suite.NotNil(err6)
	suite.Nil(report6)
}
