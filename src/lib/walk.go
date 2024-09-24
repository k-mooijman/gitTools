package lib

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetGitRepos(rootPath string, repos *Repos) {
	var gitRepos []string

	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				gitPath, _ := strings.CutSuffix(path, ".git")
				repos.addByPath(gitPath)
				gitRepos = append(gitRepos, gitPath)
				return filepath.SkipDir
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("-> variable - err = %v is of type %T \n", err, err)
	}

}

func getGitStatus(path string) string {
	//git status --porcelain
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response

}
