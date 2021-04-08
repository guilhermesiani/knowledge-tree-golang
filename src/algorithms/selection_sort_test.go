package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert.Equal(
		t,
		[]int{1, 3, 5, 8, 9},
		selectionSort([]int{8, 5, 3, 9, 1}),
	)
}

func TestSelectionSortWithNegativeValue(t *testing.T) {
	assert.Equal(
		t,
		[]int{-3, 1, 5, 7, 8},
		selectionSort([]int{5, 8, 1, 7, -3}),
	)
}
