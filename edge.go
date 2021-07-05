package reggiexpress

type Edge struct {
	wildcard bool
	pattern  string
	dest     *Node
}

func (e *Edge) Test(s string) (bool, string, *Node) {
	if len(s) < len(e.pattern) {
		return false, "", nil
	}
	if e.wildcard {
		return true, s[1:], e.dest
	} else {
		return s[:len(e.pattern)] == e.pattern, s[len(e.pattern):], e.dest
	}
}

func NewEdge(wildcard bool, pattern string, dest *Node) *Edge {
	return &Edge{
		wildcard,
		pattern,
		dest,
	}
}
