package astutil

import (
	"go/ast"
	"go/token"
)

type complexityVisitor struct {
	value int
}

func (v *complexityVisitor) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.CaseClause, *ast.CommClause, *ast.ForStmt, *ast.FuncDecl, *ast.IfStmt, *ast.RangeStmt:
		v.value++
	case *ast.BinaryExpr:
		if node.Op == token.LAND || node.Op == token.LOR {
			v.value++
		}
	}
	return v
}

// FunctionComplexity computes and returns the cyclomatic complexity
// of the given function.
func FunctionComplexity(decl *ast.FuncDecl) int {
	visitor := complexityVisitor{}
	ast.Walk(&visitor, decl)
	return visitor.value
}
