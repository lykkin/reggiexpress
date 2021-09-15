package reggiexpress

type group struct {
	// name string;
	captured bool
	// value string;
	graph flowGraph
	queue []node
}

func newGroup(queue []node, captured bool) group {
	return group{
		captured,
		newFlowGraph(),
		queue,
	}
}
