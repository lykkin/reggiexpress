package reggiexpress_test

import (
	"reggiexpress"
	"testing"
)

func TestAddEdge(t *testing.T) {
	n := reggiexpress.NewNode()
	term := reggiexpress.NewNode()
	e := reggiexpress.NewEdge(false, "asdf", term)
	n.AddEdge(e, false)
}
