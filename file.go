package main

import (
	"go/ast"
	"go/token"
)

type File struct {
	Functions []*Function `json:"functions"`
	Path      string      `json:"path"`
}

func (f *File) processFunction(decl *ast.FuncDecl, fileSet *token.FileSet) {
	function := newFromAST(decl, fileSet)
	f.Functions = append(f.Functions, function)
}
