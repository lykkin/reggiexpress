package reggiexpress

import (
	"testing"
)

func TestFlowGraphBuild(t *testing.T) {
	pattern := "asdf"
	fg := newFlowGraph()
	err, tokenStream := createTokenizer(pattern)
	if err != nil {
		t.Fail()
	}
	fg.build(tokenStream)
}

func TestFlowGraphBuildEmpty(t *testing.T) {
	pattern := ""
	fg := newFlowGraph()
	err, tokenStream := createTokenizer(pattern)
	if err != nil {
		t.Fail()
	}
	fg.build(tokenStream)
}

func TestFlowGraphProcess(t *testing.T) {
	pattern := "asdf"
	fg := newFlowGraph()
	err, tokenStream := createTokenizer(pattern)
	if err != nil {
		t.Fail()
	}
	fg.build(tokenStream)
	err, res := fg.process("asdf")
	if err != nil {
		t.Errorf("Got an error while processing: %s", err)
	}

	if len(res) < 1 {
		t.Errorf("Expected to get results, got none!")
	}
}
