package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testGraph struct {
	edges map[int][]int
}

func newTestGraph() *testGraph {
	return &testGraph{
		edges: make(map[int][]int),
	}
}

func (g *testGraph) Count() int {
	return len(g.edges)
}

func (g *testGraph) EdgesFrom(i int) []int {
	return g.edges[i]
}

func (g *testGraph) AddEdge(from int, to int) {
	if _, ok := g.edges[from]; !ok {
		g.edges[from] = []int{to}
	} else {
		g.edges[from] = append(g.edges[from], to)
	}

	// if to is not part of known nodes, add it.
	if _, ok := g.edges[to]; !ok {
		g.edges[to] = nil
	}
}

func TestGraphWithCycles(t *testing.T) {
	t.Run("cycle to itself", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 0)

		ok, cycle := VerifyAcyclic(g, 0)
		assert.False(t, ok)
		verifyPath(t, []int{0, 0}, cycle)
	})

	t.Run("cycle between 3 nodes", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		// 0->1->2, no cycles yet
		ok, _ := VerifyAcyclic(g, 0)
		assert.True(t, ok)

		g.AddEdge(2, 0)
		ok, _ = VerifyAcyclic(g, 0)
		assert.False(t, ok)
	})
}

func TestGraphWithoutCycles(t *testing.T) {
	t.Run("no cycle 1", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		ok, _ := VerifyAcyclic(g, 0)
		assert.True(t, ok)
	})

	t.Run("no cycle 2", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)
		ok, _ := VerifyAcyclic(g, 0)
		assert.True(t, ok)
	})

	t.Run("no cycle 3", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)
		g.AddEdge(0, 4)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(1, 4)
		g.AddEdge(2, 3)
		g.AddEdge(2, 4)
		g.AddEdge(3, 4)
		ok, _ := VerifyAcyclic(g, 0)
		assert.True(t, ok)
	})
}

func verifyPath(t *testing.T, expected []int, actual []int) {
	assert.Equal(t, len(expected), len(actual), "expected the length of cycle to be the same")

	for i := 0; i < len(actual); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}
