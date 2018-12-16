package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
)

type Repo struct {
	Commits []*Commit `json:"commits"`
}

func (r *Repo) processCommit(gitCommit, nextGitCommit *object.Commit) error {
	// Update the last commit container with additions and deletions
	if nextGitCommit != nil {
		additions := 0
		deletions := 0

		patch, err := gitCommit.Patch(nextGitCommit)
		if err != nil {
			return err
		}
		for _, stat := range patch.Stats() {
			additions = additions + stat.Addition
			deletions = deletions + stat.Deletion
		}

		lastCommit := r.Commits[len(r.Commits)-1]
		lastCommit.Additions = additions
		lastCommit.Deletions = deletions
	}
	nextGitCommit = gitCommit

	commit := &Commit{
		Author:  gitCommit.Author.Name,
		Email:   gitCommit.Author.Email,
		Epoch:   gitCommit.Author.When.Unix(),
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
