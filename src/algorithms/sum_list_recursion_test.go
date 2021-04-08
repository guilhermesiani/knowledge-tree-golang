package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumListRecursion(t *testing.T) {
	assert.Equal(
		t,
		26,
		sumListRecursion(&[]int{8, 5, 3, 9, 1}),
	)
	assert.Equal(
		t,
		12,
		sumListRecursion(&[]int{3, 4, 2, 1, 2}),
	)
}
