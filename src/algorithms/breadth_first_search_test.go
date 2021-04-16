package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var friends = map[string][]string{}
	friends["mine"] = []string{"romulo", "vinicius", "jhozefem"}
	friends["romulo"] = []string{"kaio", "victor"}
	friends["vinicius"] = []string{}
	friends["jhozefem"] = []string{"gisele"}
	assert.Equal(t, "gisele", breadthFirstSearch(friends))
}
