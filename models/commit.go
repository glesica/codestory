package models

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"strings"
)

type Commit struct {
	Additions int             `json:"additions"`
	Author    string          `json:"author"`
	Deletions int             `json:"deletions"`
	Email     string          `json:"email"`
	Epoch     int64           `json:"epoch"`
	Hash      string          `json:"hash"`
	Message   string          `json:"message"`
	Pkgs      map[string]*Pkg `json:"packages"`
}

func (c *Commit) processFile(gitFile *object.File) error {
	if c.Pkgs == nil {
		c.Pkgs = make(map[string]*Pkg)
	}

	if !strings.HasSuffix(gitFile.Name, ".go") {
		return nil
	}

	if strings.HasPrefix(gitFile.Name, "vendor/") {
		return nil
	}

	fileSet := token.NewFileSet()
	contents, err := gitFile.Contents()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read gitFile")
		return err
	}

	astFile, err := parser.ParseFile(fileSet, gitFile.Name, contents, 0)
	if err != nil {
		// TODO: Handle syntax errors gracefully
		// Failed to build AST
		// cmd/bindingsTranspiler/testAssets/testFailingGoCode.go:5:31: expected ';', found '{'
		fmt.Fprintln(os.Stderr, "Failed to build AST")
		return err
	}

	pkgName := astFile.Name.Name

	pkg, ok := c.Pkgs[pkgName]
	if !ok {
		pkg = &Pkg{
			Name: pkgName,
		}
		c.Pkgs[pkgName] = pkg
	}

	for _, decl := range astFile.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			pkg.processFunction(funcDecl, gitFile.Name, fileSet)
		}
	}

	return nil
}
