package awsbilling

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

func (suite *CalculatorTestSuite) TestGenerateBillingReport() {

	items := billingItemsForTest()

	report, err := generateBillingReport(items)
	suite.Nil(err)
	suite.NotNil(report)

	report2, err2 := generateBillingReport([]billingItem{})
	suite.NotNil(err2)
	suite.Nil(report2)
}
