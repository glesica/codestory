package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"strings"
)

type Repo struct {
	Commits []*Commit `json:"commits"`
}

func (r *Repo) processCommit(gitCommit *object.Commit) error {
	additions := 0
	deletions := 0
	hasParent := false
	// If the commit has a parent commit, we can retrieve the size
	// of the diff directly from the patch information. Otherwise,
	// we'll back-fill the additions for a commit with no parent
	// later on.
	stats, err := gitCommit.Stats()
	if err == nil {
		hasParent = true
		for _, stat := range stats {
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

	files, err := gitCommit.Files()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve files")
		return err
	}

	// TODO: Generalize file filter instead of doing it inline in two places
	return files.ForEach(func(file *object.File) error {
		if !strings.HasSuffix(file.Name, ".go") {
			return nil
		}
		if strings.HasPrefix(file.Name, "vendor/") {
			return nil
		}
		if !hasParent {
			lines, err := file.Lines()
			if err != nil {
				return err
			}
			commit.Additions = commit.Additions + len(lines)
		}
		return commit.processFile(file)
	})
}
