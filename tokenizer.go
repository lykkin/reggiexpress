package reggiexpress

type patternToken struct {
	isOptional bool
	pattern    string
}

type controlToken struct {
	isOptional  bool
	instruction controlSymbol
}

// ???: is there a more restrictive type we can use here?
type tokenStream chan interface{}

func createTokenizer(input string) (error, tokenStream) {
	tokenStream := make(tokenStream)
	go func() {
		lastIdx := 0

		for i := range input {
			c := input[i]
			if ct, isControl := controlSymbols[c]; isControl {
				if lastIdx != i {
					tokenStream <- patternToken{false, input[lastIdx:i]}
					lastIdx = i + 1
				}
				tokenStream <- controlToken{false, ct}
			}
		}

		if lastIdx != len(input)+1 {
			tokenStream <- patternToken{false, input[lastIdx:]}
		}

		close(tokenStream)
	}()
	return nil, tokenStream
}
