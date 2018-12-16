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

// FunctionLength computes and returns the length, in lines, of a
// function, excluding comments.
func FunctionLength(decl *ast.FuncDecl, fileSet *token.FileSet) int {
	pos := fileSet.PositionFor(decl.Pos(), false)
	end := fileSet.PositionFor(decl.End(), false)

	return end.Line - pos.Line
}
