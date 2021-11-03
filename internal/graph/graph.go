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

// Graph is a Dig-agnostic representation of a directed acyclic graph
// Nodes are identified by their position inside an ordered list
type Graph interface {
	// Count reports the total number of nodes in the graph.
	Count() int

	// Returns the indices of nodes that node i points to.
	EdgesFrom(i int) []int
}

// HasCycle checks whether the given graph contains a cycle when traversing
// the graph via DFS, starting from n-th node.
// If a cycle is detected, a slice containing the nodes in the cycle is returned.
func HasCycle(g Graph, n int) (bool, []int) {
	visited := make(map[int]interface{})
	onStack := make(map[int]bool)
	hc, cycle := hasCycleHelp(g, n, visited, onStack)
	return hc, cycle
}

func hasCycleHelp(g Graph, n int, visited map[int]interface{}, onStack map[int]bool) (bool, []int) {
	visited[n] = struct{}{}
	onStack[n] = true

	for _, neighbor := range g.EdgesFrom(n) {
		if _, ok := visited[neighbor]; !ok {
			if hasCycleHelp(g, neighbor, visited, onStack) {
				return true
			}
		} else if os, ok := onStack[neighbor]; os && ok {
			return true
		}
	}
	onStack[n] = false
	return false
}
