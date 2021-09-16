package reggiexpress

import (
	"fmt"
	"strings"
)

type edge struct {
	wildcard bool
	pattern  string
	dest     *node
}

func (e *edge) test(s string) (bool, string, *node) {
	if len(s) < len(e.pattern) {
		return false, "", nil
	}
	if e.wildcard {
		return true, s[1:], e.dest
	}
	return s[:len(e.pattern)] == e.pattern, s[len(e.pattern):], e.dest
}

func newEdge(wildcard bool, pattern string, dest *node) edge {
	return edge{
		wildcard,
		pattern,
		dest,
	}
}

func (e *edge) print(indent int) {
	buffer := "  "
	fmt.Println(strings.Repeat(buffer, indent) + e.pattern)
	e.dest.print(indent + 1)
}
