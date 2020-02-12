package chunnu_compiler

import "log"

type visitor map[string]func(n *node, p node)

func transformer(asTree ast) ast {
	nast := ast{
		kind: "Program",
		body: []node{},
	}

	asTree.context = &nast.body

	traverser(asTree, map[string]func(n *node, p node){
		"NumberLiteral": func(thisNode *node, parent node) {
			*parent.context = append(*parent.context, node{
				kind:  "NumberLiteral",
				value: thisNode.value,
			})
		},

		"CallExpression" : func(thisNode *node, parent node) {
			expression := node{
				kind: "CallExpression",
				callee: &node{
					kind: "Identifier",
					name: thisNode.name,
				},
				arguments: new([]node),
			}

			// referencing old abstract syntax tree context to new abstract syntax tree's expression arguments
			// arguments will be pushed via this context reference
			thisNode.context = expression.arguments

			if parent.kind == "CallExpression"{
				expressionStatement:= node{
					kind:"ExpressionStatement",
					expression:&expression,
				}
				// assign new node i.e expression statement to parent's context which
				// in turn refers to New Abstract Syntax tree body
				*parent.context = append(*parent.context,expressionStatement)
			}else {
				*parent.context = append(*parent.context,expression)
			}
		},
	})
	return nast
}



func traverser(asTree ast, vTool visitor) {
	traverseNode(node(asTree), node{},vTool)
}

func traverseArray(a []node, p node, v visitor) {
	for _, child := range a {
		traverseNode(child, p, v)
	}
}

func traverseNode(node, parent node, vTool visitor) {

	// applying the appropriate method of transformation method
	for key, transformMethod := range vTool {
		if key == node.kind {
			transformMethod(&node, parent)
		}
	}

	switch node.kind {
	case "Program":
		traverseArray(node.body, node, vTool)
		break
	case "CallExpression":
		traverseArray(node.params,node,vTool)
		break
	case "NumberLiteral":
		break
	default:
		log.Fatal("UnRegistered kind: "+node.kind)
	}
}
