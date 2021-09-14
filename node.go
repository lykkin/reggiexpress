package reggiexpress

import "fmt"

var id int

type Node struct {
	edges []Edge
  id int
  visitedInPrint bool
}

func NewNode() Node {
  id += 1
	return Node{
		[]Edge{},
    id,
    false,
	}
}

func (n *Node) AddEdge(e Edge, optional bool) {
	n.edges = append(n.edges, e)
}

func (n *Node) Equal(other Node) bool {
  return n.id == other.id
}

func (n *Node) Traverse(input string) []GraphNode {
	res := []GraphNode{}
	for _, e := range n.edges {
		matched, rem, next := e.Test(input)
		if matched {
      fmt.Println(rem, &next)
			res = append(res, GraphNode{
				next,
				rem,
			})
		}
	}
	return res
}

func (n *Node) Print(indent int) {
  if (n.visitedInPrint) {
    return
  }
  n.visitedInPrint = true
  fmt.Println(id)
  for _, e := range n.edges {
    e.Print(indent)
  }
  n.visitedInPrint = false
}
