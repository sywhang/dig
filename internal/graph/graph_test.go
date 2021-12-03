// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestGraph struct {
	Nodes map[int][]int
}

func newTestGraph() *TestGraph {
	return &TestGraph{
		Nodes: make(map[int][]int),
	}
}

func (g TestGraph) Order() int {
	return len(g.Nodes)
}

func (g TestGraph) EdgesFrom(u int) []int {
	return g.Nodes[u]
}

func TestGraphIsAcyclic(t *testing.T) {
	testCases := []struct {
		edges [][]int
	}{
		// 1
		{
			edges: [][]int{
				nil,
			},
		},
		// 1 --> 2 --> 3
		{
			edges: [][]int{
				{2},
				{3},
				nil,
			},
		},
		// 1 ---> 2 -------> 3
		// |                 ^
		// ------------------|
		{
			edges: [][]int{
				{2, 3},
				{3},
				nil,
			},
		},
		// 1 --> 2 --> 3    5 --> 6
		// |           ^    ^
		// |-----------|    |
		// ----------> 4 ----
		{
			edges: [][]int{
				{2, 3, 4},
				{3},
				nil,
				{5},
				{6},
				nil,
			},
		},
	}
	for _, tt := range testCases {
		g := newTestGraph()
		for i, neighbors := range tt.edges {
			g.Nodes[i+1] = neighbors
		}
		ok, _ := IsAcyclic(g)
		assert.True(t, ok)
	}
}

func TestGraphIsCyclic(t *testing.T) {
	testCases := []struct {
		edges [][]int
		cycle []int
	}{
		//
		// 0 ---> 1 ---> 2
		// ^             |
		// |-------------|
		{
			edges: [][]int{
				{1},
				{2},
				{3},
				{0},
			},
			cycle: []int{0, 1, 2, 3},
		},
		//
		// 0 ---> 1 ---> 2
		//        ^      |
		//        |------|
		{
			edges: [][]int{
				{1},
				{2},
				{1},
			},
			cycle: []int{1, 2},
		},
		//
		// 0 ---> 1 ---> 2 ----> 3
		// |      ^      |       ^
		// |      |------|       |
		// -----------------------
		{
			edges: [][]int{
				{1, 3},
				{2},
				{1, 3},
				nil,
			},
			cycle: []int{1, 2},
		},
	}
	for _, tt := range testCases {
		g := newTestGraph()
		for i, neighbors := range tt.edges {
			g.Nodes[i] = neighbors
		}
		ok, c := IsAcyclic(g)
		assert.False(t, ok)
		for _, node := range tt.cycle {
			assert.Contains(t, c, node)
		}
	}
}
