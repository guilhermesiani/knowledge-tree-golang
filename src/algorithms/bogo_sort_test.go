package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBogoSort(t *testing.T) {
	test := [5]int{1,2,3,4,5}
	bogoSort(&test)
	assert.IsType(t, [5]int{}, test)
	assert.NotEqual(t, [5]int{1,2,3,4,5}, test)
}