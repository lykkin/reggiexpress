package reggiexpress_test

import (
	"reggiexpress"
	"testing"
)

//func TestFlowGraphBuild(t *testing.T) {
//	pattern := "asdf"
//	fg := reggiexpress.NewFlowGraph()
//  fg.Build(pattern)
//  fg.Print()
//}
//
//func TestFlowGraphBuildEmpty(t *testing.T) {
//	pattern := ""
//	fg := reggiexpress.NewFlowGraph()
//  fg.Build(pattern)
//}

func TestFlowGraphProcess(t *testing.T) {
	pattern := "asdf"
	fg := reggiexpress.NewFlowGraph()
  fg.Build(pattern)
  err, res := fg.Process("asdf")
  if err != nil {
    t.Errorf("Got an error while processing: %s", err)
  }

  if len(res) < 1 {
    t.Errorf("Expected to get results, got none!")
  }
}
