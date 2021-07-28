package reggiexpress

type Group struct {
	// name string;
	captured bool
	// value string;
	graph FlowGraph
	queue []Node
}

func NewGroup(queue []Node, captured bool) Group {
	return Group{
		captured,
		NewFlowGraph(),
		queue,
	}
}
