package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
)

type Repo struct {
	Commits []*Commit `json:"commits"`
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
