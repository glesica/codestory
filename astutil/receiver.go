package astutil

import "go/ast"

// ReceiverName computes a readable string name for a method receiver.
func ReceiverName(expr ast.Expr) string {
	name := ""
	switch t := expr.(type) {
	case *ast.Ident:
		name = t.Name
	case *ast.StarExpr:
		name = "*" + ReceiverName(t.X)
	}
	return name
}
