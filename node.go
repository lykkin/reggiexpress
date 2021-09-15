package reggiexpress

import "fmt"

var id int

type node struct {
	edges []edge
  id int
  visitedInPrint bool
}

func newNode() node {
  id += 1
	return node{
		[]edge{},
    id,
    false,
	}
}

func (n *node) addEdge(e edge, optional bool) {
	n.edges = append(n.edges, e)
}

func (n *node) traverse(input string) []graphNode {
	res := []graphNode{}
	for _, e := range n.edges {
		matched, rem, next := e.test(input)
		if matched {
			res = append(res, graphNode{
				next,
				rem,
			})
		}
	}
	return res
}

func (n *node) print(indent int) {
  if (n.visitedInPrint) {
    return
  }
  n.visitedInPrint = true
  fmt.Println("-> ", n.id)
  for _, e := range n.edges {
    e.print(indent)
  }
  n.visitedInPrint = false
}
