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

// struct for keeping track of useful info needed for cycle detection.
type cycleDetectInfo struct {
	visited   map[int]interface{}
	onStack   map[int]bool
	backtrack map[int]int
}

// VerifyAcyclic checks whether the given graph contains a cycle when traversing
// the graph via DFS, starting from n-th node.
// If a cycle is detected, a slice containing the nodes in the cycle is returned.
func VerifyAcyclic(g Graph, n int) (bool, []int) {
	info := &cycleDetectInfo{
		visited:   make(map[int]interface{}),
		onStack:   make(map[int]bool),
		backtrack: make(map[int]int),
	}
	ok := isAcyclic(g, n, info)
	cycle := getCyclePath(n, info.backtrack)
	return ok, cycle
}

func isAcyclic(g Graph, n int, info *cycleDetectInfo) bool {
	info.visited[n] = struct{}{}
	info.onStack[n] = true

	for _, neighbor := range g.EdgesFrom(n) {
		info.backtrack[neighbor] = n
		if _, ok := info.visited[neighbor]; !ok {
			if !isAcyclic(g, neighbor, info) {
				return false
			}
		} else if os, ok := info.onStack[neighbor]; os && ok {
			return false
		}
	}
	info.onStack[n] = false
	return true
}

func getCyclePath(start int, bt map[int]int) []int {
	path := []int{start}

	for curr, ok := bt[start]; ok && curr != start; curr = bt[curr] {
		path = append([]int{curr}, path...)
	}
	path = append([]int{start}, path...)
	return path
}
