package chunnu_compiler

func Compile(input string)string{
	tokens := tokenizer(input)
	ast:= parser(tokens)
	nast:= transformer(ast)
	output := codegen(node(nast))
	return output
}
