package main

import (
	"ChunnuCompiler/chunnu_compiler"
	"fmt"
)

func main(){
	program := "(add 10 (subtract 10 6))"
	out := chunnu_compiler.Compile(program)
	fmt.Println(out)
}