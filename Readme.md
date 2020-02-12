 ## Chunnu Compiler
 
 #### Description
 
 This compiler transforms expression from `lisp syntax` to `C syntax`
 It is just project to understand the working of compiler phases i.e 
 - **Lexical analysis** (Breakage into tokens)
 - **Syntactical analysis** (Creation of Abstract Syntax tree)
 - **Transformation** (Creation of Transformed Syntax tree according to platform/requirements) 
 - **Code Generation** (Code generated according to new transformed syntax tree)
 
  Each phase is coded in different `go file` for easy understanding
  
 ### Usage
 Input: `(add 2 (subtract 10 5))`
 
 Output: `add(2, subtract(10, 5));` 
 ## Credits
 This compiler is made with help of [hazbo's tiny compiler](https://github.com/hazbo/the-super-tiny-compiler)
 (or you could say re-implementation of it :P ) for study and getting experience of **Go lang and Compiler**