package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapReduceTestSuit struct {
	suite.Suite
}

func (suit *MapReduceTestSuit) Test1Row() {
	build()
	// assert.Equal(suit.T(), "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
}

func TestMapReduceTestSuit(t *testing.T) {
	suite.Run(t, new(MapReduceTestSuit))
}
