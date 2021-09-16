package reggiexpress

import (
	"testing"
)

func testBuild(t *testing.T, pattern string) flowGraph {
	fg := newFlowGraph()
	err, tokenStream := createTokenizer(pattern)
	if err != nil {
		t.Errorf("Error from tokenizer creation: %s", err)
	}
	err = fg.build(tokenStream)
	if err != nil {
		t.Errorf("Error from build: %s", err)
	}
	return fg
}

func TestFlowGraphSimple(t *testing.T) {
	fg := testBuild(t, "asdf")
	err, res := fg.process("asdf")
	if err != nil {
		t.Errorf("Got an error while processing: %s", err)
	}

	if len(res) < 1 {
		t.Errorf("Expected to get results, got none!")
	}
}

func TestFlowGraphEmpty(t *testing.T) {
	testBuild(t, "")
}

func TestFlowGraphGroups(t *testing.T) {
	fg := testBuild(t, "asdf(1234)qwerty")
	err, res := fg.process("asdf1234qwerty")
	if err != nil {
		t.Errorf("Got an error while processing: %s", err)
	}

	if len(res) < 1 {
		t.Errorf("Expected to get results, got none!")
	}
}

/* BENCHMARKS */

var result interface{}

func BenchmarkBuildGroups(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fg := newFlowGraph()
		_, tokenStream := createTokenizer("asdf(1234)(qwerty)")
		fg.build(tokenStream)
		result = fg
	}
}
