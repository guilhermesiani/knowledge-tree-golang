package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBogoSort(t *testing.T) {
	test := [5]int{4,5,6,1,7}
	assert.Equal(t, [4]int{4,5,6,1}, test)
}