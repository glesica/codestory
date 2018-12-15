package astutil

import (
	"go/ast"
)

type panicVisitor struct {
	count int
}

func (v *panicVisitor) Visit(n ast.Node) ast.Visitor {
	return v
}

// FunctionPanics counts the number of calls to `panic` that occur
// inside the given function.
func FunctionPanics(decl *ast.FuncDecl) int {
	return 0
}
