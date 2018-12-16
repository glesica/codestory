package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
)

type Repo struct {
	Commits []*Commit `json:"commits"`
}

func (r *Repo) processCommit(gitCommit *object.Commit, prevGitCommit *object.Commit) error {
	additions := 0
	deletions := 0
	if prevGitCommit != nil {
		patch, err := prevGitCommit.Patch(gitCommit)
		if err != nil {
			return err
		}
		for _, stat := range patch.Stats() {
			additions = additions + stat.Addition
			deletions = deletions + stat.Deletion
		}
	}
	commit := &Commit{
		Additions: additions,
		Author:    gitCommit.Author.Name,
		Deletions: deletions,
		Email:     gitCommit.Author.Email,
		Epoch:     gitCommit.Author.When.Unix(),
		Hash:      gitCommit.Hash.String(),
		Message:   gitCommit.Message,
	}

	r.Commits = append(r.Commits, commit)

	tree, err := gitCommit.Tree()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve tree")
		return err
	}

	return tree.Files().ForEach(commit.processFile)
}
