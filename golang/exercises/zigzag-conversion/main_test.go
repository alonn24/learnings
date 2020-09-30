package zigzag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ZigZagTestSuit struct {
	suite.Suite
}

func (suit *ZigZagTestSuit) Test1Row() {
	assert.Equal(suit.T(), "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
}

func (suit *ZigZagTestSuit) Test2Row() {
	assert.Equal(suit.T(), "PYAIHRNAPLSIIG", convert("PAYPALISHIRING", 2))
}

func (suit *ZigZagTestSuit) Test3Rows() {
	assert.Equal(suit.T(), "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}

func (suit *ZigZagTestSuit) Test4Rows() {
	assert.Equal(suit.T(), "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
}

func (suit *ZigZagTestSuit) TestMoreRowsThanLetters() {
	assert.Equal(suit.T(), "PAYPALISHIRING", convert("PAYPALISHIRING", 99))
}

func TestZigZagTestSuit(t *testing.T) {
	suite.Run(t, new(ZigZagTestSuit))
}
