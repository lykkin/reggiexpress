package reggiexpress

import "fmt"

type FlowGraph struct {
	Source Node
	End    Node
}

func NewFlowGraph() FlowGraph {
  s := NewNode()
  e := NewNode()
	return FlowGraph{
    s,
    e,
	}
}

type GraphNode struct {
	node  Node
	input string
}

func (fg *FlowGraph) Process(input string) (error, []string) {
	stack := []GraphNode{
		GraphNode{
			fg.Source,
			input,
		},
	}
	res := []string{}
  i := 0
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		n := stack[lastIdx]
    stack = stack[:lastIdx]
    i += 1
		//if &n.node == &fg.End {
    fmt.Print(n.node.id, fg.End.id)
    if n.node.Equal(fg.End) {
      fmt.Print(1)
			res = append(res, input[len(n.input):])
		} else {
			stack = append(stack, n.node.Traverse(input)...)
		}
    fmt.Println("length", len(stack))
	}
  fmt.Println("i", i)
	return nil, res
}

func (fg *FlowGraph) Print() {
  fmt.Println("Begin")
  fg.Source.Print(0)
}

func (fg *FlowGraph) Build(pattern string) {
  currentNode := &fg.Source;
  // TODO: optimize this into a set of offsets
  subpattern := "";
  for i := range pattern {
    c := pattern[i]
    if _, isControl := ControlSymbols[c]; isControl {
      if len(subpattern) > 0 {
        n := NewNode()
        currentNode.AddEdge(
          NewEdge(false, subpattern, n),
          c == '?',
        );
        currentNode = &n
        subpattern = ""
      }
    } else {
      subpattern += string(c);
    }
  }
  if len(subpattern) > 0 {
    n := NewNode()
    currentNode.AddEdge(
      NewEdge(false, subpattern, n),
      false,
    );
    currentNode = &n
  }
  currentNode.AddEdge(
    NewEdge(false, "", fg.End),
    false,
  )
}
