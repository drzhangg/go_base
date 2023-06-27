package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

func main() {
	// Clone the given repository to the given directory
	//Info("git clone https://github.com/drzhangg/gitops")

	url := "https://github.com/drzhangg/gitops"

	tempDir := ""
	var err error

	if tempDir, err = os.MkdirTemp(os.TempDir(), ""); err != nil {
		fmt.Println("failed to create a temp directory, error is ", err)
		return
	}

	auth := &http.BasicAuth{
		Username: "drzhangg",
		Password: "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",
	}

	if _, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName("main"),
		Auth:          auth,
	}); err != nil {
		fmt.Println("clone err:",err)
		return
	}

	fmt.Println("tempDir:",tempDir)


	var repo *git.Repository

	//if repo, err = git.PlainClone(tempDir, false, &git.CloneOptions{
	if repo, err = git.PlainOpen(tempDir); err == nil {
		var wd *git.Worktree

		if wd, err = repo.Worktree(); err == nil {
			if err := wd.Checkout(&git.CheckoutOptions{
				Branch: plumbing.NewBranchReferenceName("dev2"),
				Create: true,
				Keep:   true,
			}); err != nil {
				fmt.Println("checkout error: ", err)
				return
			}

			// commit and push
			if _, err := wd.Add("."); err != nil {
				fmt.Println("add err:",err)
				return
			}
			if _, err := wd.Commit("msg", &git.CommitOptions{
				Committer: nil,
			}); err != nil {
				fmt.Println("commit err:",err)
				return
			}


		}
	}else {
		fmt.Println("PlainOpen err:",err)
	}
}
