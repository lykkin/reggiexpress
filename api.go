package reggiexpress
type regex struct {
  fg flowGraph
}

type TestResult struct {
  match bool
  groups []string
  //TODO: named groups
}

func (r *regex) Test(input string) (error, TestResult) {
  err, matches := r.fg.process(input)
  if err != nil {
    // TODO: handle processing errors
  }
  return nil, TestResult {
    len(matches) > 0,
    matches,
  }
}

func NewRegex(pattern string) regex {
  err, ts := createTokenizer(pattern)
  if err != nil {
    // TODO: handle tokenizer errors
  }
  fg := newFlowGraph()
  fg.build(ts)
  return regex{fg}
}

