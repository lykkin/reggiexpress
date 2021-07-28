package reggiexpress_test

import (
	"reggiexpress"
	"testing"
)

func TestEdgeTest(t *testing.T) {
	matches := []string{
		"asdf",
		"asdf1",
		"asdfasdf",
	}
	nonMatches := []string{
		"",
		"sdfg",
		"1asdf",
		"vcvcv",
		"asd",
		"asdd",
	}
	pattern := "asdf"
	e := reggiexpress.NewEdge(false, pattern, reggiexpress.Node{})
	for _, m := range matches {
		doesMatch, rest, _ := e.Test(m)
		if !doesMatch || rest != m[len(pattern):] {
			t.Fail()
		}
	}
	for _, m := range nonMatches {
		doesMatch, rest, _ := e.Test(m)
		if doesMatch || len(rest) > 0 && rest != m[len(pattern):] {
			t.Fail()
		}
	}
}

func TestEdgeWildcardTest(t *testing.T) {
	pattern := "."
	e := reggiexpress.NewEdge(true, pattern, reggiexpress.Node{})
	doesMatch, rest, _ := e.Test("asdf")
	if !doesMatch || rest != "sdf" {
		t.Fail()
	}
	doesMatch, rest, _ = e.Test("")
	if doesMatch || rest != "" {
		t.Fail()
	}
}

func TestEdgeEmptyPattern(t *testing.T) {
	pattern := ""
	e := reggiexpress.NewEdge(false, pattern, reggiexpress.Node{})
	doesMatch, rest, _ := e.Test("")
	doesMatch, rest, _ = e.Test("")
	if !doesMatch || rest != "" {
		t.Fail()
	}
}
