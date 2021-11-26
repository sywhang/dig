// Copyright (c) 2019 Uber Technologies, Inc.
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
	// Copy over all the providers
	providersCopy := make(map[key][]*constructorNode, len(s.providers))
	nodesCopy := make([]*constructorNode, len(s.nodes))
	valuesCopy := make(map[key]reflect.Value, len(s.values))
	groupsCopy := make(map[key][]reflect.Value, len(groups))

	return &Scope{
		parentScope: parent,
	}
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
	} else {
		return nil
	}
}
