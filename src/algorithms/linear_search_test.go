package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearchSuccess(t *testing.T) {
	assert.Equal(
		t,
		3,
		linearSearch([5]string{"a", "r", "g", "q", "l"}, "q"),
	)
	assert.Equal(
		t,
		1,
		linearSearch([5]string{"a", "r", "g", "q", "l"}, "r"),
	)
}

func TestLinearSearchFail(t *testing.T) {
	assert.Equal(
		t,
		-1,
		linearSearch([5]string{"a", "r", "g", "q", "l"}, "m"),
	)
}