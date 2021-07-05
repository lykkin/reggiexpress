package reggiexpress

type Engine struct {
	queue []*Node
	graph *FlowNetwork
}

func NewEngine() *Engine {
	q := []*Node{}
	return &Engine{
		q,
		NewFlowNetwork(),
	}
}

func (e *Engine) Build(pattern string) {
}
