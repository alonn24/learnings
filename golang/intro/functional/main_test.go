package functional

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FunctionalTestSuit struct {
	suite.Suite
}

func (suit *FunctionalTestSuit) Test1Row() {
	run()
	// assert.Equal(suit.T(), "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
}

func TestFunctionalTestSuit(t *testing.T) {
	suite.Run(t, new(FunctionalTestSuit))
}
