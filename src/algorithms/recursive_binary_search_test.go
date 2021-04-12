package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecursiveBinarySearchSuccess(t *testing.T) {
	binarySearchResult := recursiveBinarySearch([]int{0, 1, 3, 4}, 3)
	assert.Equal(t, 2, binarySearchResult)
}

func TestRecursiveBinarySearchOnBigSliceSuccess(t *testing.T) {
	binarySearchResult := recursiveBinarySearch([]int{0, 1, 3, 5, 7, 9, 13, 23, 25, 32}, 23)
	assert.Equal(t, 7, binarySearchResult)
}

func TestRecursiveBinarySearchOnBigSliceSuccessSearchingHalfDown(t *testing.T) {
	binarySearchResult := recursiveBinarySearch([]int{0, 1, 3, 5, 7, 9, 13, 23, 25, 32}, 3)
	assert.Equal(t, 2, binarySearchResult)
}

func TestRecursiveBinarySearchFail(t *testing.T) {
	binarySearchResult := recursiveBinarySearch([]int{0, 1, 3, 4}, 2)
	assert.Equal(t, -1, binarySearchResult)
}
