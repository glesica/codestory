package models

import (
	"go/ast"
	"go/token"
)

type Pkg struct {
	Funs    []*Fun    `json:"funs"`
	Name    string    `json:"name"`
	Structs []*Struct `json:"structs"`
}

func (p *Pkg) processFunction(decl *ast.FuncDecl, fileName string, fileSet *token.FileSet) {
	fun := newFromAST(decl, fileName, fileSet)
	p.Funs = append(p.Funs, fun)
}
