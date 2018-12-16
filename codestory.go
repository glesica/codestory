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

	logOptions := &git.LogOptions{
		From:  gitHead.Hash(),
		Order: git.LogOrderCommitterTime,
	}
	commitIter, err := gitRepo.Log(logOptions)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve log")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	repo := Repo{}

	var nextCommit *object.Commit
	err = commitIter.ForEach(func(commit *object.Commit) error {
		err := repo.processCommit(commit, nextCommit)
		nextCommit = commit
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
