package reggiexpress

import (
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
	e := newEdge(false, pattern, nil)
	for _, m := range matches {
		doesMatch, rest, _ := e.test(m)
		if !doesMatch || rest != m[len(pattern):] {
			t.Fail()
		}
	}
	for _, m := range nonMatches {
		doesMatch, rest, _ := e.test(m)
		if doesMatch || len(rest) > 0 && rest != m[len(pattern):] {
			t.Fail()
		}
	}
}

func TestEdgeWildcardTest(t *testing.T) {
	pattern := "."
	e := newEdge(true, pattern, nil)
	doesMatch, rest, _ := e.test("asdf")
	if !doesMatch || rest != "sdf" {
		t.Fail()
	}
	doesMatch, rest, _ = e.test("")
	if doesMatch || rest != "" {
		t.Fail()
	}
}

func TestEdgeEmptyPattern(t *testing.T) {
	pattern := ""
	e := newEdge(false, pattern, nil)
	doesMatch, rest, _ := e.test("")
	doesMatch, rest, _ = e.test("")
	if !doesMatch || rest != "" {
		t.Fail()
	}
}
