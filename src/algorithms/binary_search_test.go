package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchSuccess(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 4}, 3)
	assert.Equal(t, 2, binarySearchResult)
}

func TestBinarySearchFail(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 4}, 2)
	assert.Equal(t, -1, binarySearchResult)
}
