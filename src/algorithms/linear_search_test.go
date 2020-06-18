package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	assert.Equal(
		t,
		3,
		linearSearch([5]string{"a", "r", "g", "q", "l"}, "q"),
	)
}