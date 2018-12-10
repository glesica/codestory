package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"strings"
)

func main() {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-git",
	})
	if err != nil {
		panic("ugh")
	}

	head, err := repo.Head()
	if err != nil {
		panic("argh")
	}

	logIter, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		panic("whatever")
	}

	c, err := logIter.Next()
	if err != nil {
		panic("barf")
	}

	tree, err := c.Tree()
	if err != nil {
		panic("stupid")
	}

	maxDeclsName := ""
	maxDecls := 0

	treeIter:= tree.Files()
	treeIter.ForEach(func(file *object.File) error {
		if !strings.HasSuffix(file.Name, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		src, err := file.Contents()
		if err != nil {
			panic("stupid source")
		}

		parsedFile, err := parser.ParseFile(fset, file.Name, src, 0)
		if err != nil {
			panic(err)
		}

		//fmt.Print(file.Name)
		//fmt.Print(" - ")
		//fmt.Println(len(parsedFile.Decls))

		thisLen := len(parsedFile.Decls)
		if thisLen > maxDecls {
			maxDeclsName = file.Name
			maxDecls = thisLen
		}

		return nil
	})

	fmt.Print(maxDeclsName)
	fmt.Print(" - ")
	fmt.Println(maxDecls)

	//
	//
	//for _, entry := range tree.Entries {
	//	fmt.Println(entry.Mode)
	//	//fmt.Println(entry.Name)
	//}
	//
	//logIter.ForEach(func(c *object.Commit) error {
	//	//fmt.Println(c.Hash)
	//	//fmt.Println(c.Message)
	//	_, err := c.Tree()
	//	if err != nil {
	//		panic("stupid")
	//	}
	//
	//	return nil
	//})
}
