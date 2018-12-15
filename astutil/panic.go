package astutil

import (
	"go/ast"
)

type panicVisitor struct {
	count int
}

func (v *panicVisitor) Visit(node ast.Node) ast.Visitor {
	if call, ok := node.(*ast.CallExpr); ok {
		fun := call.Fun
		if ident, ok := fun.(*ast.Ident); ok {
			if ident.Name == "panic" {
				v.count++
			}
		}
	}
	return v
}

// FunctionPanics counts the number of calls to `panic` that occur
// inside the given function.
func FunctionPanics(decl *ast.FuncDecl) int {
	visitor := panicVisitor{}
	ast.Walk(&visitor, decl)
	return visitor.count
}
