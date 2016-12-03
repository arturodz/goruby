package main

import (
	"fmt"
	"strconv"
)

func main() {

}

// Token represents the value of a variable or an instruction in the
// programming language.
type Token struct {
	elem   string
	valInt int
	valStr string
	next   int
}

// tokenizer is responsible for breaking up a sentence apart into tokens.
// Basically is the Lexical Analyzer
func tokenizer(input string, position int) (token *Token, EOF bool, err error) {
	if position > len(input)-1 {
		return nil, true, nil
	}

	current_character := input[position : position+1]

	// Check if val is integer
	if intVal, err := strconv.Atoi(current_character); err == nil {
		token := Token{
			elem:   "INTEGER",
			valInt: intVal,
			next:   position + 1,
		}
		return token, false, nil
	}

	// Check if val is a plus char
	if current_character == '+' {
		token := Token{
			elem:   "PLUS",
			valStr: current_character,
			next:   position + 1,
		}
		return token, false, nil
	}

	err := raiseException("Error parsing input")
	return nil, nil, err
}

func raiseException(errVal string) {
	return errors.New(errVal)
}
