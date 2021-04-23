package algorithms

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := map[string]map[string]float64{}
	graph["start"] = map[string]float64{"a": 6, "b": 2}
	graph["a"] = map[string]float64{"end": 1}
	graph["b"] = map[string]float64{"a": 3, "end": 5}
	graph["end"] = map[string]float64{}

	costs := map[string]float64{"a": 6, "b": 2, "end": math.Inf(0)}

	parents := map[string]string{"a": "start", "b": "start", "end": ""}

	result := dijkstra(graph, costs, parents)
	assert.Equal(t, []string{"b", "a", "end"}, result)
}
