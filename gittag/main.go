package main

import (
	"fmt"
	"github.com/ForgeCloud/saas/go/common/pkg/git2go"
)

func main() {
	repo, err := git2go.NewRepoFromFilesystem(
		"/Users/david.macneil/go/src/github.com/fr-dam/various",
		true,
		git2go.RepoOpts{},
	)
	if err != nil {
		fmt.Println(err)
	}
	err = repo.Push(false)
	if err != nil {
		fmt.Println(err)
	}
}
