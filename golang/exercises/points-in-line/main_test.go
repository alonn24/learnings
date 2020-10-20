package points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	assert.Equal(t, 4, maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}

func TestExamples2(t *testing.T) {
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {1, 2}, {1, 3}, {2, 2}}))
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {2, 1}, {3, 1}, {2, 2}}))
	assert.Equal(t, 0, maxPoints([][]int{}))
	assert.Equal(t, 2, maxPoints([][]int{{0, 0}, {94911150, 94911151}, {94911151, 94911152}}))
	// assert.Equal(t, 2, maxPoints([][]int{{0, 0}, {94911151, 94911150}, {94911152, 94911151}}))
	assert.Equal(t, 4, maxPoints([][]int{{3, 1}, {12, 3}, {3, 1}, {-6, -1}}))
}
