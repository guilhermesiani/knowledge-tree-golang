package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var friends = map[string][]string{}
	friends["mine"] = []string{"romulo", "vinicius", "jhozefem"}
	friends["romulo"] = []string{"kaio", "victor"}
	friends["jhozefem"] = []string{"gisele", "vinicius"}
	assert.Equal(t, "gisele", breadthFirstSearch(friends))
}
