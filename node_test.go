package reggiexpress

import (
	"testing"
)

func TestAddEdge(t *testing.T) {
	n := newNode()
	term := newNode()
	e := newEdge(false, "asdf", &term)
	n.addEdge(e, false)
}
