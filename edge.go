package reggiexpress

import (
  "fmt"
  "strings"
)

type Edge struct {
	wildcard bool
	pattern  string
	dest     Node
}

func (e *Edge) Test(s string) (bool, string, Node) {
	if len(s) < len(e.pattern) {
		return false, "", Node{}
	}
	if e.wildcard {
		return true, s[1:], e.dest
	}
  return s[:len(e.pattern)] == e.pattern, s[len(e.pattern):], e.dest
}

func NewEdge(wildcard bool, pattern string, dest Node) Edge {
	return Edge{
		wildcard,
		pattern,
		dest,
	}
}

func (e *Edge) Print(indent int) {
  buffer := "  "
  fmt.Println(strings.Repeat(buffer, indent) + e.pattern)
  e.dest.Print(indent + 1)
}
