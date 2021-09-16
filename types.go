package reggiexpress

type controlSymbol int

const (
  startGroup controlSymbol = iota
  endGroup
)

var controlSymbols  = map[byte]controlSymbol {
  '(': startGroup,
  ')': endGroup,
}
