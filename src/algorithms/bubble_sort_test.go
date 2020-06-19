package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	assert.Equal(
		t,
		[5]int{1,3,5,8,9},
		bubbleSort([5]int{8,5,3,9,1}),
	)
	assert.Equal(
		t,
		[5]int{-3,1,5,7,8},
		bubbleSort([5]int{5,8,1,7,-3}),
	)
}