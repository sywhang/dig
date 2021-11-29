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

package dig

import "reflect"

// newScope creates an empty new Scope.
func (s *Scope) cloneScope() *Scope {
	// TODO (sungyoon): It may be possible to do a "lazy"
	// resolution of the graph by keeping just deltas and
	// applying it when it's needed, rather than dumbly
	// copying over everything at once.

	child := &Scope{
		parentScope: s,
		providers:   make(map[key][]*constructorNode),
		nodes:       make([]*constructorNode, 0),
		values:      make(map[key]reflect.Value),
		groups:      make(map[key][]reflect.Value),
		invokerFn:   s.invokerFn,
	}

	// child should hold a separate graph holder
	gh := &graphHolder{
		orders: make(map[key]int),
		s:      child,
	}

	child.gh = gh

	return child
}

// Scope creates a new empty Scope from the
// current Scope's context.
func (s *Scope) Scope(opts ...ScopeOption) *Scope {
	newS := s.cloneScope()
	s.childScopes = append(s.childScopes, newS)
	return newS
}

// GetScopesUntilRoot creates a list of Scopes
// have to traverse through until the current node.
func (s *Scope) GetScopesUntilRoot() []*Scope {
	if s.parentScope != nil {
		return append(s.parentScope.GetScopesUntilRoot(), s)
	}
	return []*Scope{s}
}
