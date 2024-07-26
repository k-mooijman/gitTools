package lib

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetGitRepos(rootPath string, repos *Repos) []string {
	var gitRepos []string

	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				gitPath, _ := strings.CutSuffix(path, ".git")
				//fmt.Printf(" .................... \n\n")

				//fmt.Printf("%v\n - Remote: %v \n - Current Branch: %v  \n", path, getGitRemote(gitPath), getGitCurrentBranch(gitPath))
				//fmt.Printf(" - Status: \n%v ", getGitStatus(gitPath))
				//branches := getGitGetBranches(gitPath)

				//for i := 0; i < len(branches); i++ {
				//	fmt.Println(branches[i])
				//}
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

	return gitRepos
}

//ls-remote --exit-code --heads origin test

//func getGitRemote(path string) string {
//
//	cmd := exec.Command("git", "remote", "-v")
//	cmd.Dir = path
//	cmdResponse, _ := cmd.Output()
//
//	return cleanRemoteToOne(string(cmdResponse))
//}

//func getGitGetBranches(path string) []string {
//
//	cmd := exec.Command("git", "branch", "-a")
//	cmd.Dir = path
//	cmdResponse, _ := cmd.Output()
//	response := string(cmdResponse)
//
//	response, _ = strings.CutSuffix(response, "\n")
//	temp := strings.Split(response, "\n")
//
//	return temp
//
//}

func getGitStatus(path string) string {
	//git status --porcelain
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response

}

//func getGitCurrentBranch(path string) string {
//
//	cmd := exec.Command("git", "branch", "--show-current")
//	cmd.Dir = path
//	cmdResponse, _ := cmd.Output()
//	response := string(cmdResponse)
//	response, _ = strings.CutSuffix(response, "\n")
//
//	return response
//
//}

}

func cleanRemoteToOne(response string) string {

	response = strings.Replace(response, "(fetch)", "", -1)
	response = strings.Replace(response, "(push)", "", -1)
	response, _ = strings.CutSuffix(response, "\n")

	temp := strings.Split(response, "\n")
	if len(temp) < 2 {
		return ""
	}
	if temp[0] == temp[1] {
		response, _ = strings.CutPrefix(temp[0], "origin")
		response = strings.TrimSpace(response)
		return response
	} else {
		return ""
	}
}
//func cleanRemoteToOne(response string) string {
//
//	response = strings.Replace(response, "(fetch)", "", -1)
//	response = strings.Replace(response, "(push)", "", -1)
//	response, _ = strings.CutSuffix(response, "\n")
//
//	temp := strings.Split(response, "\n")
//
//	if temp[0] == temp[1] {
//		response, _ = strings.CutPrefix(temp[0], "origin")
//		response = strings.TrimSpace(response)
//		return response
//	} else {
//		return ""
//	}
//}
