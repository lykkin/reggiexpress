package reggiexpress

type Group struct {
	// name string;
	captured bool
	// value string;
	graph *FlowNetwork
	queue []*Node
}

func NewGroup(queue []*Node, captured bool) *Group {
	return &Group{
		captured,
		NewFlowNetwork(),
		queue,
	}
}
