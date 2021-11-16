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

import (
	"bytes"
	"fmt"

	"go.uber.org/dig/internal/digreflect"
)

type cycleEntry struct {
	Key  key
	Func *digreflect.Func
}

type errCycleDetected struct {
	Path []cycleEntry
}

func (e errCycleDetected) Error() string {
	// We get something like,
	//
	//   foo provided by "path/to/package".NewFoo (path/to/file.go:42)
	//   	depends on bar provided by "another/package".NewBar (somefile.go:1)
	//   	depends on baz provided by "somepackage".NewBar (anotherfile.go:2)
	//   	depends on foo provided by "path/to/package".NewFoo (path/to/file.go:42)
	//
	b := new(bytes.Buffer)

	for i, entry := range e.Path {
		if i > 0 {
			b.WriteString("\n\tdepends on ")
		}
		fmt.Fprintf(b, "%v provided by %v", entry.Key, entry.Func)
	}
	return b.String()
}

// IsCycleDetected returns a boolean as to whether the provided error indicates
// a cycle was detected in the container graph.
func IsCycleDetected(err error) bool {
	_, ok := RootCause(err).(errCycleDetected)
	return ok
}
