package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AtoITestSuit struct {
	suite.Suite
}

func (suit *AtoITestSuit) TestEmpty() {
	assert.Equal(suit.T(), 0, myAtoi(""))
	assert.Equal(suit.T(), 0, myAtoi("asdwq"))
	assert.Equal(suit.T(), 49, myAtoi("49"))
	assert.Equal(suit.T(), 49, myAtoi("49asdqw"))
	assert.Equal(suit.T(), -49, myAtoi("-49"))
	assert.Equal(suit.T(), 0, myAtoi("    -+-49asd"))
	assert.Equal(suit.T(), 0, myAtoi("    -dqw49asd"))
	assert.Equal(suit.T(), 0, myAtoi("dqw49asd"))
}

func (suit *AtoITestSuit) TestEmpty2() {
	assert.Equal(suit.T(), -2147483648, myAtoi("-91283472332"))
	assert.Equal(suit.T(), 2147483647, myAtoi("21474836460"))
	assert.Equal(suit.T(), -1, myAtoi("-000000000000000000000000000000000000000000000000001"))
	assert.Equal(suit.T(), 2147483647, myAtoi("20000000000000000000000000000000000000000000000000"))
}

func TestAtoITestSuit(t *testing.T) {
	suite.Run(t, new(AtoITestSuit))
}
