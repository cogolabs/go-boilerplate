package examplepkg

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var testVal = "FOOBAR"

type exampletypeSuite struct {
	suite.Suite
	et *ExampleType
}

func (s *exampletypeSuite) SetupSuite() {
	// Run before the first test
	s.et = New(testVal)
}

func (s *exampletypeSuite) TearDownSuite() {
	// Run after the last test
}

func (s *exampletypeSuite) SetupTest() {
	// Run before each test
}

func (s *exampletypeSuite) TearDownTest() {
	// Run after each test
}

func (s *exampletypeSuite) TestMyFunc() {
	// If the following statement fails, the test will end immediately
	s.Require().Equal(s.et.PublicVal, testVal)

	tests := []struct {
		n        int
		expected string
	}{
		{1, "f"},
		{3, "foo"},
		{1000, "foobar"},
	}

	var val string
	for _, tt := range tests {
		val = s.et.MyFunc(tt.n)
		// If the following statement fails, the test will fail, but continue
		s.Assert().Equal(val, tt.expected)
	}
}

func TestMyType(t *testing.T) {
	tests := new(exampletypeSuite)
	suite.Run(t, tests)
}
