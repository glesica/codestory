package main

import "go/ast"

type File struct {
	Functions []*Function `json:"functions"`
	Path      string      `json:"path"`
}

func (f *File) processFunction(declFunc *ast.FuncDecl) {
	function := &Function{
		Arity: declFunc.Type.Params.NumFields(),
		Name:  declFunc.Name.String(),
	}

	f.Functions = append(f.Functions, function)
}
