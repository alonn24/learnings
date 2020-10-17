package validnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValidNumberTestSuit struct {
	suite.Suite
}

func (suit *ValidNumberTestSuit) TestValid() {
	assert.Equal(suit.T(), true, isNumber("0"))
	assert.Equal(suit.T(), true, isNumber(" 0.1 "))
	assert.Equal(suit.T(), true, isNumber(".1"))
	assert.Equal(suit.T(), true, isNumber("3."))
	assert.Equal(suit.T(), true, isNumber("2e10"))
	assert.Equal(suit.T(), true, isNumber(" -90e3   "))
	assert.Equal(suit.T(), true, isNumber(" 6e-1"))
	assert.Equal(suit.T(), true, isNumber("53.5e93"))
	assert.Equal(suit.T(), true, isNumber("2e10"))
	assert.Equal(suit.T(), true, isNumber("    0000000000000000000000000000000000000000000000000012e10"))
}

func (suit *ValidNumberTestSuit) TestNotValid() {
	assert.Equal(suit.T(), false, isNumber("abc"))
	assert.Equal(suit.T(), false, isNumber("1 a"))
	assert.Equal(suit.T(), false, isNumber(" 1e"))
	assert.Equal(suit.T(), false, isNumber("e3"))
	assert.Equal(suit.T(), false, isNumber(" 99e2.5 "))
	assert.Equal(suit.T(), false, isNumber(" --6 "))
	assert.Equal(suit.T(), false, isNumber("-+3"))
	assert.Equal(suit.T(), false, isNumber("95a54e53"))
	assert.Equal(suit.T(), false, isNumber("    0000000000000000000000000000000000000000000000000012e10s"))
}

func TestValidNumberTestSuit(t *testing.T) {
	suite.Run(t, new(ValidNumberTestSuit))
}
