package bhaskara

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcMaxProfit(t *testing.T) {
	assert.Equal(t, 55, CalcMaxProfit(350, 40, 5))
	assert.Equal(t, 60, CalcMaxProfit(350, 50, 5))
}