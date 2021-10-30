package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testGraph struct {
	edges map[int64][]int64
}

func newTestGraph() *testGraph {
	return &testGraph{
		edges: make(map[int64][]int64),
	}
}

func (g *testGraph) Count() int {
	return len(g.edges)
}

func (g *testGraph) EdgesFrom(i int64) []int64 {
	return g.edges[i]
}

func (g *testGraph) AddEdge(from int64, to int64) {
	if _, ok := g.edges[from]; !ok {
		g.edges[from] = []int64{to}
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

		assert.True(t, HasCycle(g, 0))
	})

	t.Run("cycle between 3 nodes", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		// 0->1->2, no cycles yet
		assert.False(t, HasCycle(g, 0))

		g.AddEdge(2, 0)
		assert.True(t, HasCycle(g, 0))
		assert.True(t, HasCycle(g, 2))
	})
}

func TestGraphWithoutCycles(t *testing.T) {
	t.Run("no cycle 1", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		assert.False(t, HasCycle(g, 0))
	})

	t.Run("no cycle 2", func(t *testing.T) {
		g := newTestGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)
		assert.False(t, HasCycle(g, 0))
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
		assert.False(t, HasCycle(g, 0))
		assert.False(t, HasCycle(g, 1))
		assert.False(t, HasCycle(g, 2))
		assert.False(t, HasCycle(g, 3))
		assert.False(t, HasCycle(g, 4))
	})
}
