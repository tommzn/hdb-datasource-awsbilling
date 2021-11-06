package awsbilling

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReaderTestSuite struct {
	suite.Suite
}

func TestReaderTestSuite(t *testing.T) {
	suite.Run(t, new(ReaderTestSuite))
}

func (suite *ReaderTestSuite) TestReadS3ObjectContent() {

	objectKey1 := "fixtures/BillingReport.csv"
	content1, err1 := readFixtures(objectKey1)
	suite.Nil(err1)
	lines1, err1 := readObject(objectKey1, content1, true)
	suite.Nil(err1)
	suite.Len(lines1, 6)

	objectKey2 := "fixtures/BillingReport.csv.gz"
	content2, err2 := readFixtures(objectKey2)
	suite.Nil(err2)
	lines2, err2 := readObject(objectKey2, content2, true)
	suite.Nil(err2)
	suite.Len(lines2, 6)

	lines3, err3 := readObject(objectKey2, content1, true)
	suite.NotNil(err3)
	suite.Len(lines3, 0)

	emptyFile := "fixtures/empty.BillingReport.csv"
	content4, err4 := readFixtures(emptyFile)
	suite.Nil(err4)
	lines4, err4 := readObject(emptyFile, content4, true)
	suite.Nil(err4)
	suite.Len(lines4, 1)
}
