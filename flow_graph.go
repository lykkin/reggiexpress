package reggiexpress

import (
  "errors"
)

type flowGraph struct {
	Source node
	End    node
}

func newFlowGraph() flowGraph {
  s := newNode()
  e := newNode()
	return flowGraph{
    s,
    e,
	}
}

type graphNode struct {
	node  *node
	input string
}

func (fg *flowGraph) process(input string) (error, []string) {
	stack := []graphNode{
		graphNode{
			&fg.Source,
			input,
		},
	}
	res := []string{}
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		n := stack[lastIdx]
    stack = stack[:lastIdx]
		if n.node == &fg.End {
			res = append(res, input[len(n.input):])
		} else {
			stack = append(stack, n.node.traverse(input)...)
		}
	}
  if len(res) == 0 {
    return errors.New("no matches found"), res
  }
  return nil, res
}

func (fg *flowGraph) print() {
  fg.Source.print(0)
}

func (fg *flowGraph) build(tokens tokenStream) {
  // TODO: what happens when built twice?
  currentNode := &fg.Source;
  for token := range tokens {
    if token.isControl {
      // This is always false for now
    } else {
      n := newNode()
      currentNode.addEdge(
        newEdge(false, token.pattern, &n),
        token.isOptional,
      );
      currentNode = &n
    }
  }
  currentNode.addEdge(
    newEdge(false, "", &fg.End),
    false,
  )
}
