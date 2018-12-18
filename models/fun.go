package models

import (
	"github.com/glesica/codestory/astutil"
	"go/ast"
	"go/token"
)

type Fun struct {
	Arity      int    `json:"arity"`
	Complexity int    `json:"complexity"`
	File       string `json:"file"`
	Length     int    `json:"length"`
	Name       string `json:"name"`
	Panics     int    `json:"panics"`
}

func newFromAST(decl *ast.FuncDecl, fileName string, fileSet *token.FileSet) *Fun {
	return &Fun{
		Arity:      astutil.FunctionArity(decl),
		Complexity: astutil.FunctionComplexity(decl),
		File:       fileName,
		Length:     astutil.FunctionLength(decl, fileSet),
		Name:       astutil.FunctionName(decl),
		Panics:     astutil.FunctionPanics(decl),
	}
}
