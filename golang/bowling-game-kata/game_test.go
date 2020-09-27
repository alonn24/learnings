package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BowlingTestSuite struct {
	suite.Suite
}

func (suite *BowlingTestSuite) SetupTest() {
	// setup code
}

func (suite *BowlingTestSuite) TestCreation() {
	game := Game{}
	assert.Equal(suite.T(), game.score(), 0)
}

func (suite *BowlingTestSuite) TestUncompletedRoll() {
	game := Game{}
	game.roll(5)
	assert.Equal(suite.T(), game.score(), 0)
}

func (suite *BowlingTestSuite) TestTwoRolls() {
	game := Game{}
	game.roll(2).roll(3)
	assert.Equal(suite.T(), game.score(), 2+3)
}

func (suite *BowlingTestSuite) TestTwoTurns() {
	game := Game{}
	game.roll(2).roll(3)
	game.roll(2).roll(3)
	assert.Equal(suite.T(), game.score(), 5+5)
}

func (suite *BowlingTestSuite) TestSpare() {
	game := Game{}
	game.roll(5).roll(5)
	game.roll(2)
	assert.Equal(suite.T(), game.score(), 10+2)
}

func (suite *BowlingTestSuite) TestStrike() {
	game := Game{}
	game.roll(10)
	game.roll(2).roll(3)
	assert.Equal(suite.T(), game.score(), 10+2+3)
}

func TestBowlingTestSuite(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}
