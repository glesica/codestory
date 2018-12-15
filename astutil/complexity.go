package astutil

import (
	"go/ast"
	"go/token"
)

type complexityVisitor struct {
	value int
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.CaseClause, *ast.CommClause, *ast.ForStmt, *ast.FuncDecl, *ast.IfStmt, *ast.RangeStmt:
		v.value++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.value++
		}
	}
	return v
}

func FunctionComplexity(decl *ast.FuncDecl) int {
	visitor := complexityVisitor{}
	ast.Walk(&visitor, decl)
	return visitor.value
}
