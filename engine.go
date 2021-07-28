package reggiexpress

type Engine struct {
	queue []Node
	graph FlowGraph
}

func NewEngine() Engine {
	q := []Node{}
	return Engine{
		q,
		NewFlowGraph(),
	}
}

func (e *Engine) Build(pattern string) {
}
