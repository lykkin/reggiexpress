package reggiexpress

type token struct {
  isOptional bool
  isControl bool
  pattern string
}

type tokenStream chan token

func createTokenizer(input string) (error, chan token) {
  tokenStream := make(tokenStream)
  go func(){
    lastIdx := 0
    for i := range input {
      c := input[i]
      if _, isControl := ControlSymbols[c]; isControl {
        if lastIdx != i {
          tokenStream <- token {false, false, input[lastIdx:i]}
          lastIdx = i
        }
        tokenStream <- token {false, true, string(c)}
      }
    }
    if lastIdx != len(input) + 1 {
      tokenStream <- token {false, false, input[lastIdx:]}
    }
    close(tokenStream)
  }()
  return nil, tokenStream
}
