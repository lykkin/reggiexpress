package reggiexpress

import (
	"testing"
)

func TestCreateTokenizer(t *testing.T) {
  err, tokenStream := createTokenizer("asdf")
  if err != nil {
    t.Fail()
  }
  tokens := make([]token, 0)

  for token := range tokenStream {
    tokens = append(tokens, token)
  }

  if len(tokens) != 1  && tokens[0].pattern != "asdf" {
    t.Fail()
  }
}
