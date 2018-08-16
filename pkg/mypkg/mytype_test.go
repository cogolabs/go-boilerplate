package mypkg

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var testVal = "FOOBAR"

type mytypeSuite struct {
	suite.Suite
	myt *MyType
}

func (s *mytypeSuite) SetupSuite() {
	// Run before the first test
	s.myt = New(testVal)
}

func (s *mytypeSuite) TearDownSuite() {
	// Run after the last test
}

func (s *mytypeSuite) SetupTest() {
	// Run before each test
}

func (s *mytypeSuite) TearDownTest() {
	// Run after each test
}

func (s *mytypeSuite) TestMyFunc() {
	s.Require().Equal(s.myt.PublicVal, testVal)
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
		val = s.myt.MyFunc(tt.n)
		s.Equal(val, tt.expected)
	}
}

func TestMyType(t *testing.T) {
	tests := new(mytypeSuite)
	suite.Run(t, tests)
}
