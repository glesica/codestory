package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
)

func main() {
	gitRepo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open current directory")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	gitHead, err := gitRepo.Head()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve HEAD")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	commitIter, err := gitRepo.Log(&git.LogOptions{From: gitHead.Hash()})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve log")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	repo := Repo{}

	var prevCommit *object.Commit
	err = commitIter.ForEach(func(commit *object.Commit) error {
		err := repo.processCommit(commit, prevCommit)
		prevCommit = commit
		return err
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	jsonBytes, err := json.Marshal(repo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to encode JSON")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Print(string(jsonBytes))
}
