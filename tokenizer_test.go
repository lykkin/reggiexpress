package reggiexpress

import (
	"reflect"
	"testing"
)

func getParameterError(parameterName string) string {
	return "Mismatched " + parameterName + ": actual %s, expected %s"
}

func checkToken(t *testing.T, actual interface{}, expected interface{}) {
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Token type mismatch, actual %s, expected %s", reflect.TypeOf(actual), reflect.TypeOf(expected))
	}
	switch cast := actual.(type) {
	case patternToken:
		if cast.isOptional != expected.(patternToken).isOptional {
			t.Errorf(getParameterError("isOptional"), actual, expected)
		}
		if cast.pattern != expected.(patternToken).pattern {
			t.Errorf(getParameterError("pattern"), actual, expected)
		}
	case controlToken:
		if cast.isOptional != expected.(controlToken).isOptional {
			t.Errorf(getParameterError("isOptional"), actual, expected)
		}
		if cast.instruction != expected.(controlToken).instruction {
			t.Errorf(getParameterError("instruction"), actual, expected)
		}
	default:
		t.Errorf("Unknown token type %s", reflect.TypeOf(actual))
	}
}

func testTokens(t *testing.T, pattern string, expectedTokens []interface{}) {
	err, tokenStream := createTokenizer(pattern)
	if err != nil {
		t.Fail()
	}
	tokens := make([]interface{}, 0)

	for token := range tokenStream {
		tokens = append(tokens, token)
	}

	if len(tokens) != len(expectedTokens) {
		t.Fail()
	}

	for i := range expectedTokens {
		checkToken(t, tokens[i], expectedTokens[i])
	}

}

func TestCreateTokenizer(t *testing.T) {
	testTokens(t, "asdf", []interface{}{
		patternToken{false, "asdf"},
	})
}

func TestTokenizeGroup(t *testing.T) {
	testTokens(t, "asdf(1234)qwerty", []interface{}{
		patternToken{false, "asdf"},
		controlToken{false, startGroup},
		patternToken{false, "1234"},
		controlToken{false, endGroup},
		patternToken{false, "qwerty"},
	})
}
