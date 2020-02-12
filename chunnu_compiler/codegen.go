package chunnu_compiler

import (
	"log"
	"strings"
)

func codegen(n node) string  {
	switch n.kind {
	case "Program":
		var nodeArr []string
		for _, currNode := range n.body{
			nodeArr = append(nodeArr,codegen(currNode))
		}
		// Separating each node of body by next line
		return strings.Join(nodeArr,"\n")
	case "ExpressionStatement":
		return codegen(*n.expression) + ";"

	case "CallExpression":
		var argumentsArr []string
		callExp := codegen(*n.callee)

		for _, currNode := range *n.arguments{
			argumentsArr = append(argumentsArr,codegen(currNode))
		}
		argumentsStr := strings.Join(argumentsArr,",")
		return callExp + "("+argumentsStr+")"
	case "Identifier":
		return n.name

	case "NumberLiteral":
		return n.value

	default:
		log.Fatal("Unable to recognise this node "+n.kind)
		return ""
	}
}
