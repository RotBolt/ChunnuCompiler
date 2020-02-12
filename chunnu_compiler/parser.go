package chunnu_compiler

import (
	"log"
)

var pc int
var pt []token
func parser(tokens []token) ast {
	pc := 0
	pt =tokens
	ast := ast{
		kind: "Program",
		body: []node{},
	}
	for pc < len(tokens) {
		ast.body = append(ast.body, walk(&pc))
	}
	return ast
}

func walk(pc *int) node {

	token := pt[*pc]

	if token.category == "number_literal" {
		*pc++
		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}

	if token.category == "parenthesis" && token.value == "(" {
		// increment and check for other tokens, "(" will not be added to ast
		*pc++
		token :=pt[*pc]

		thisNode := node{
			kind: "CallExpression",
			name : token.value,
			params:[]node{},
		}

		*pc++
		token = pt[*pc]

		for token.category != "parenthesis" ||
			(token.category == "parenthesis" && token.value != ")"){
			thisNode.params = append(thisNode.params,walk(pc))
			token = pt[*pc]
		}

		*pc++

		return thisNode
	}
	log.Fatal("Unable to recognize "+token.category)
	return node{}

}
