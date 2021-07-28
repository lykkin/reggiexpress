package reggiexpress

type FlowGraph struct {
	Source *Node
	End    *Node
}

func NewFlowGraph() FlowGraph {
	r := NewNode()
	return FlowGraph{
		&r,
		&r,
	}
}

type GraphNode struct {
	node  Node
	input string
}

func (fg *FlowGraph) Process(input string) (error, []string) {
	stack := []GraphNode{
		GraphNode{
			*fg.Source,
			input,
		},
	}
	res := []string{}
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		n := stack[lastIdx]
		if &n.node == fg.End {
			res = append(res, input[len(n.input):])
		} else {
			stack = append(stack[:lastIdx], n.node.Traverse(input)...)
		}
	}
	return nil, res
}

func (fg *FlowGraph) Build(pattern string) {
	//for i, c := range pattern
}
