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
}