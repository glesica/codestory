package main

import (
	"github.com/glesica/codestory/astutil"
	"go/ast"
	"go/token"
)

type Function struct {
	Arity      int    `json:"arity"`
	Complexity int    `json:"complexity"`
	Length     int    `json:"length"`
	Name       string `json:"name"`
	Panics     int    `json:"panics"`
}

func newFromAST(decl *ast.FuncDecl, fileSet *token.FileSet) *Function {
	return &Function{
		Arity:      astutil.FunctionArity(decl),
		Complexity: astutil.FunctionComplexity(decl),
		Length:     astutil.FunctionLength(decl, fileSet),
		Name:       astutil.FunctionName(decl),
		Panics:     astutil.FunctionPanics(decl),
	}
}
