package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindHighestValueOnSlice(t *testing.T) {
	assert.Equal(t, 27, findHighestValue([]int{9, 5, 1, 7, 12, 27, 18, 15}))
}
