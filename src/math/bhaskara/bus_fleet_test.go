package bhaskara

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcMaxProfit(t *testing.T) {
	assert.Equal(
		t,
		55,
		calcMaxProfit(
			350,
			40,
			5,
		),
	)
}