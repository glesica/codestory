package models

import (
	"go/ast"
	"go/token"
)

type Pkg struct {
	Funs    map[string]*Fun    `json:"funs"`
	Name    string             `json:"name"`
	Structs map[string]*Struct `json:"structs"`
}

func (p *Pkg) processFunction(decl *ast.FuncDecl, fileName string, fileSet *token.FileSet) {
	if p.Funs == nil {
		p.Funs = make(map[string]*Fun)
	}

	if p.Structs == nil {
		p.Structs = make(map[string]*Struct)
	}

	fun := newFromAST(decl, fileName, fileSet)
	p.Funs[fun.Name] = fun
}
