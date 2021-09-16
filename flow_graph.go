package reggiexpress

import (
	"errors"
)

type flowGraph struct {
	Source node
	End    node
}

func newFlowGraph() flowGraph {
	return flowGraph{
		newNode(),
		newNode(),
	}
}

type graphNode struct {
	// TODO: axe all the pointers
	node     *node
	fragment string
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
		if n.node.id == fg.End.id {
			res = append(res, input[len(n.fragment):])
		} else {
			stack = append(stack, n.node.traverse(n.fragment)...)
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

func (fg *flowGraph) build(tokens tokenStream) error {
	// TODO: what happens when built twice?
	currentNode := &fg.Source
tokenIterator:
	for it := range tokens {
		switch token := it.(type) {
		case patternToken:
			n := newNode()
			currentNode.addEdge(
				newEdge(false, token.pattern, &n),
				token.isOptional,
			)
			currentNode = &n

		case controlToken:
			// TODO: break this out
			switch it {
			case startGroup:
				subFg := newFlowGraph()
				currentNode.addEdge(
					newEdge(false, "", &subFg.Source),
					token.isOptional,
				)
				subFg.build(tokens)
				currentNode = &subFg.End

			case endGroup:
				break tokenIterator
			}

		default:
			return errors.New("unknown token type encountered")
		}
	}
	currentNode.addEdge(
		newEdge(false, "", &fg.End),
		false,
	)
	return nil
}
