package greetings

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GreetingsTestSuite struct {
	suite.Suite
	name string
}

func (suite *GreetingsTestSuite) SetupTest() {
	// setup code
}

func (suite *GreetingsTestSuite) TestHello() {
	msg, err := Hello("alony")
	assert.Regexp(suite.T(), regexp.MustCompile(`\balony\b`), msg)
	assert.Nil(suite.T(), err)
}

func (suite *GreetingsTestSuite) TestHelloEmpty() {
	msg, err := Hello("")
	assert.Equal(suite.T(), msg, "")
	assert.NotNil(suite.T(), err)
}

func (suite *GreetingsTestSuite) TestHellos() {
	ans, err := Hellos("alony")
	assert.Len(suite.T(), ans, 1)
	assert.Regexp(suite.T(), regexp.MustCompile(`\balony\b`), ans["alony"])
	assert.Nil(suite.T(), err)
}

func TestGreetingsTestSuite(t *testing.T) {
	suite.Run(t, new(GreetingsTestSuite))
}
