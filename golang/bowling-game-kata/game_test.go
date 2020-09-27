package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BowlingTestSuite struct {
	suite.Suite
}

const PERFECT_ROUNDS = 240
const PERFECT_GAME = 300

func (suite *BowlingTestSuite) SetupTest() {
	// setup code
}

func (suite *BowlingTestSuite) TestCreation() {
	game := Game{}
	assert.Equal(suite.T(), 0, game.score())
}

func (suite *BowlingTestSuite) TestUncompletedRoll() {
	game := Game{}
	game.roll(5)
	assert.Equal(suite.T(), 5, game.score())
}

func (suite *BowlingTestSuite) TestTwoRolls() {
	game := Game{}
	game.roll(2).roll(3)
	assert.Equal(suite.T(), 5, game.score())
}

func (suite *BowlingTestSuite) TestTwoTurns() {
	game := Game{}
	game.roll(2).roll(3)
	game.roll(2).roll(3)
	assert.Equal(suite.T(), 10, game.score())
}

func (suite *BowlingTestSuite) TestSpare() {
	game := Game{}
	game.roll(5).roll(5)
	game.roll(2)
	assert.Equal(suite.T(), 14, game.score())
}

func (suite *BowlingTestSuite) TestStrike() {
	game := Game{}
	game.roll(10)
	game.roll(2).roll(3)
	assert.Equal(suite.T(), 20, game.score())
}

func (suite *BowlingTestSuite) Test5Strikes() {
	game := Game{}
	game.roll(10).roll(10).roll(10).roll(10).roll(10)
	assert.Equal(suite.T(), 120, game.score())
}

func (suite *BowlingTestSuite) TestTwoStrikesAndRoll() {
	game := Game{}
	game.roll(10).roll(10).roll(5)
	assert.Equal(suite.T(), 45, game.score())
}

func (suite *BowlingTestSuite) TestPerfectGame() {
	game := Game{}
	for i := 0; i < 12; i++ {
		game.roll(10)
	}
	assert.Equal(suite.T(), PERFECT_GAME, game.score())
}

func (suite *BowlingTestSuite) TestZeroLastRound() {
	game := Game{}
	for i := 0; i < 9; i++ {
		game.roll(10)
	}
	game.roll(0).roll(0).roll(0)
	assert.Equal(suite.T(), PERFECT_ROUNDS, game.score())
}

func (suite *BowlingTestSuite) TestLastRoundSpare() {
	game := Game{}
	for i := 0; i < 9; i++ {
		game.roll(10)
	}
	game.roll(5).roll(5).roll(5)
	assert.Equal(suite.T(), PERFECT_ROUNDS+15+15, game.score())
}

func TestBowlingTestSuite(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}
