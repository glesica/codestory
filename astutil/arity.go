package astutil

import "go/ast"

// FunctionArity computes and returns the arity (number of parameters)
// of the given function.
func FunctionArity(decl *ast.FuncDecl) int {
	return decl.Type.Params.NumFields()
}
