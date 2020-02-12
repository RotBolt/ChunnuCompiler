package ChunnuCompiler

import "log"

var pc int
func parser(tokens []token) ast {
	pc := 0

	ast := ast{
		kind: "Program",
		body: []node{},
	}
	for pc < len(tokens) {
		ast.body = append(ast.body, walk(tokens))
	}
	return ast
}

func walk(tokens []token) node {
	token := tokens[pc]

	if token.category == "number_literal" {
		pc++
		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}

	if token.category == "parenthesis" && token.value == "(" {
		// increament and check for other tokens, "(" will not be added to ast
		pc++
		token :=tokens[pc]

		thisNode := node{
			kind: "CallExpression",
			name : token.value,
			params:[]node{},
		}

		pc++
		token = tokens[pc]

		for token.category != "parenthesis" ||
			(token.category == "parenthesis" && token.value != ")"){
			thisNode.params = append(thisNode.params,walk(tokens))
			token = tokens[pc]
		}

		pc++

		return thisNode
	}
	log.Fatal("Unable to recognize "+token.category)
	return node{}

}
