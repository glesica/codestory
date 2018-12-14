package main

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
	Files   []*File
	Hash    string
	Message string
}

func (c *Commit) processFile(gitFile *object.File) error {
	if !strings.HasSuffix(gitFile.Name, ".go") {
		return nil
	}

	if strings.HasPrefix(gitFile.Name, "vendor") {
		return nil
	}

	file := &File{
		Path: gitFile.Name,
	}

	c.Files = append(c.Files, file)

	fileSet := token.NewFileSet()
	contents, err := gitFile.Contents()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read gitFile")
		return err
	}

	astFile, err := parser.ParseFile(fileSet, gitFile.Name, contents, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to build AST")
		return err
	}

	for _, decl := range astFile.Decls {
		if declFunc, ok := decl.(*ast.FuncDecl); ok {
			file.processFunction(declFunc)
		}
	}

	return nil
}

type File struct {
	Functions []*Function
	Path      string
}

func (f *File) processFunction(declFunc *ast.FuncDecl) {
	function := &Function{
		Arity: declFunc.Type.Params.NumFields(),
		Name:  declFunc.Name.String(),
	}

	f.Functions = append(f.Functions, function)
}

type Function struct {
	Arity      int
	Complexity int
	Lines      int
	Name       string
	Panics     int
}

type Repo struct {
	Commits []*Commit
}

func (r *Repo) processCommit(gitCommit *object.Commit) error {
	commit := &Commit{
		Hash:    gitCommit.Hash.String(),
		Message: gitCommit.Message,
	}

	r.Commits = append(r.Commits, commit)

	tree, err := gitCommit.Tree()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve tree")
		return err
	}

	return tree.Files().ForEach(commit.processFile)
}
