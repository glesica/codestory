package astutil

import (
	"go/ast"
	"go/token"
)

// FunctionName builds a readable string name for the given function.
func FunctionName(decl *ast.FuncDecl) string {
	name := ""
	if decl.Recv != nil && decl.Recv.NumFields() > 0 {
		name = name + "." + ReceiverName(decl.Recv.List[0].Type)
	}
	return name + decl.Name.Name
}

func FunctionLength(decl *ast.FuncDecl, fileSet *token.FileSet) int {
	pos := fileSet.Position(decl.Pos())
	end := fileSet.Position(decl.End())

	return end.Line - pos.Line
}
