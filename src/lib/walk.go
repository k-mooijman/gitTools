package lib

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetGitRepos(rootPath string) []string {
	var gitRepos []string

	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				gitPath, _ := strings.CutSuffix(path, ".git")

				fmt.Printf("%v - %v - %v  \n", path, getGitRemote(gitPath), getGitCurrentBranch(gitPath))

				gitRepos = append(gitRepos, gitPath)
				return filepath.SkipDir
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("-> variable - err = %v is of type %T \n", err, err)
	}

	return gitRepos
}

func getGitRemote(path string) string {

	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()

	response := cleanRemoteToOne(string(cmdResponse))

	return response

}

func getGitCurrentBranch(path string) string {

	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response

}

func cleanRemoteToOne(response string) string {

	response = strings.Replace(response, "(fetch)", "", -1)
	response = strings.Replace(response, "(push)", "", -1)
	response, _ = strings.CutSuffix(response, "\n")

	temp := strings.Split(response, "\n")

	if temp[0] == temp[1] {
		response, _ = strings.CutPrefix(temp[0], "origin")
		response = strings.TrimSpace(response)
		return response
	} else {
		return ""
	}
}
