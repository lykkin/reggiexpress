package reggiexpress

type FlowNetwork struct {
	Source *Node
	End    *Node
}

func NewFlowNetwork() *FlowNetwork {
	r := NewNode()
	return &FlowNetwork{
		r,
		r,
	}
}

type NetworkNode struct {
	node  *Node
	input string
}

func (fg *FlowNetwork) Process(input string) (error, []string) {
	stack := []NetworkNode{
		NetworkNode{
			fg.Source,
			input,
		},
	}
	res := []string{}
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		n := stack[lastIdx]
		if n.node == fg.End {
			res = append(res, input[len(n.input):])
		} else {
			stack = append(stack[:lastIdx], n.node.Traverse(input)...)
		}
	}
	return nil, res
}

func (fg *FlowNetwork) Build(pattern string) {
	//for i, c := range pattern
}
