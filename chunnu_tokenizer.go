package ChunnuCompiler

import (
	"unicode"
)

type token struct {
	category string
	value string
}

func tokenizer(input string) []token  {
	current := 0
	var tokens []token

	// converting string to slice of characters (just to count individual characters)
	characters := []rune(input)
	for current < len(characters){

		currChar := characters[current]

		if currChar == '('{
			tokens = append(tokens,token{
				category: "parenthesis",
				value:    "(",
			})
			current++
			continue
		}

		if currChar == ')' {
			tokens = append(tokens,token{
				category: "parenthesis",
				value:    ")",
			})
			current++
			continue
		}

		// just skip it baby
		if currChar == ' '{
			current++
			continue
		}

		if unicode.IsNumber(currChar){
			numString :=""

			for unicode.IsNumber(currChar){
				numString += string(currChar)
				current++
				currChar = characters[current]
			}

			tokens = append(tokens,token{
				category: "number_literal",
				value:    numString,
			})

			continue
		}

		if unicode.IsLetter(currChar) {
			letterString := ""
			
			for unicode.IsLetter(currChar){
				letterString += string(currChar)
				current++
				currChar = characters[current]
			}
			
			tokens = append(tokens, token{
				category: "string_literal",
				value:    letterString,
			})

			continue
		}
	}
	return tokens
}