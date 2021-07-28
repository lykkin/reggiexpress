package reggiexpress

type Node struct {
	edges []Edge
}

func NewNode() Node {
	return Node{
		[]Edge{},
	}
}

func (n *Node) AddEdge(e Edge, optional bool) {
	n.edges = append(n.edges, e)
}

func (n *Node) Traverse(input string) []GraphNode {
	res := []GraphNode{}
	for _, e := range n.edges {
		matched, rem, next := e.Test(input)
		if matched {
			res = append(res, GraphNode{
				next,
				rem,
			})
		}
	}
	return res
}
