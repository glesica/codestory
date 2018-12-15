package astutil

import "go/ast"

func FunctionArity(decl *ast.FuncDecl) int {
	return decl.Type.Params.NumFields()
}
