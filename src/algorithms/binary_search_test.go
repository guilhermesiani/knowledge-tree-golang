package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchSuccess(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 4}, 3)
	assert.Equal(t, 2, binarySearchResult)
}

func TestBinarySearchOnBigSliceSuccess(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 5, 7, 9, 13, 23, 25, 32}, 23)
	assert.Equal(t, 7, binarySearchResult)
}

func TestBinarySearchOnBigSliceSuccessSearchingHalfDown(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 5, 7, 9, 13, 23, 25, 32}, 5)
	assert.Equal(t, 3, binarySearchResult)
}

func TestBinarySearchFail(t *testing.T) {
	binarySearchResult := binarySearch([]int{0, 1, 3, 4}, 2)
	assert.Equal(t, -1, binarySearchResult)
}
