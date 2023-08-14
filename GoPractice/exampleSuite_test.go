package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
}

func (es *ExampleSuite) TestTrue() {
	es.True(false)
}

func (es *ExampleSuite) TestFalse() {
	es.False(false)
}

func (es *ExampleSuite) SetupSuite() {
	es.T().Log("Inside Setup Suite")
}

func (es *ExampleSuite) TearDownSuite() {
	es.T().Log("Inside Tear Down Setup Suite")
}

func (es *ExampleSuite) SetupTest() {
	es.T().Log("Inside Setup Test")
}

func (es *ExampleSuite) TearDownTest() {
	es.T().Log("Inside Tear Down Test")
}

func (es *ExampleSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestTrue" {
		es.T().Log("I'm only for this test")
	}
}
