package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	binarySearchResult := binarySearch(0, [4]int{4,1,0,3}, 0, 4)
	assert.Equal(t, 2, binarySearchResult)
}