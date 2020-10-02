package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckerTestSuit struct {
	suite.Suite
}

func (suit *CheckerTestSuit) TestEmpty() {
	assert.Equal(suit.T(), 6, strongPasswordChecker(""))
}

func (suit *CheckerTestSuit) TestNoChanges() {
	assert.Equal(suit.T(), 0, strongPasswordChecker("abcA12"))
	assert.Equal(suit.T(), 0, strongPasswordChecker("012345678901234567aA"))
}

func (suit *CheckerTestSuit) TestMissingSpecial() {
	assert.Equal(suit.T(), 2, strongPasswordChecker("012345"))
	assert.Equal(suit.T(), 2, strongPasswordChecker("abcdefg"))
	assert.Equal(suit.T(), 2, strongPasswordChecker("ABCDEFG"))
	assert.Equal(suit.T(), 1, strongPasswordChecker("abcDEFG"))
	assert.Equal(suit.T(), 1, strongPasswordChecker("123DEFG"))
	assert.Equal(suit.T(), 1, strongPasswordChecker("abc123"))
}

func (suit *CheckerTestSuit) TestChangeForSpecial() {
	assert.Equal(suit.T(), 2, strongPasswordChecker("aaabbb"))
	assert.Equal(suit.T(), 2, strongPasswordChecker("aaaBBB"))
}

/* INSERT */
func (suit *CheckerTestSuit) TestInsertOnly() {
	assert.Equal(suit.T(), 3, strongPasswordChecker("aA1"))
}

func (suit *CheckerTestSuit) TestInsertOverSpecial() {
	assert.Equal(suit.T(), 3, strongPasswordChecker("abc"))
}

/* DELETE */
func (suit *CheckerTestSuit) TestDelete() {
	assert.Equal(suit.T(), 2, strongPasswordChecker("aA01234567890123456789"))
}

func (suit *CheckerTestSuit) TestDeleteForRepeating() {
	assert.Equal(suit.T(), 1, strongPasswordChecker("aaabcD1"))
}

func (suit *CheckerTestSuit) TestDeleteAndChange() {
	// delete 2 and change for upper + lower case
	assert.Equal(suit.T(), 4, strongPasswordChecker("0123456789012345678912"))
}

func (suit *CheckerTestSuit) TestMultipleRepeatingWithDelete() {
	// aaAaa1aa1aa1aa1aa1aa - 6
	assert.Equal(suit.T(), 6, strongPasswordChecker("aaaaaaaaaaaaaaaaaaaa"))
}

func (suit *CheckerTestSuit) TestDeleteToFixRepeat() {
	// aa[a]bb[b]aa[a]bb[b]aa[a]bb[b]aaAAA123 - 6
	assert.Equal(suit.T(), 6, strongPasswordChecker("aaabbbaaabbbaaabbbaaAAB123"))
	// aa[a]bb[b]aa[a]bb[b]aa[a]bb[b]aaAAA111
	assert.Equal(suit.T(), 6+2, strongPasswordChecker("aaabbbaaabbbaaabbbaaAAA111"))
}

func TestCheckerTestSuit(t *testing.T) {
	suite.Run(t, new(CheckerTestSuit))
}
