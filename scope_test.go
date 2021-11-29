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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScopeTree(t *testing.T) {
	t.Parallel()
	c := New()
	s1 := c.Scope()
	s2 := c.Scope()
	s3 := s1.Scope()

	t.Run("verify Container tree", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, s1.parentScope, c.scope)
		assert.Equal(t, s2.parentScope, c.scope)

		assert.Equal(t, s3.parentScope, s1)
		assert.NotEqual(t, s3.parentScope, s2)
	})

	t.Run("GetScopesUntilRoot returns scopes in tree path in order of distance from root", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, []*Scope{c.scope, s1, s3}, s3.GetScopesUntilRoot())
		assert.Equal(t, []*Scope{c.scope, s1, s3}, s3.GetScopesUntilRoot())
	})
}
