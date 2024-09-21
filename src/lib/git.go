package lib

import (
	"os/exec"
	"strings"
)

func getGitGetBranches(path string) []string {

	cmd := exec.Command("git", "branch", "-a")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)

	response, _ = strings.CutSuffix(response, "\n")
	temp := strings.Split(response, "\n")

	return temp

}

func getGitCurrentBranch(path string) string {

	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response

}

func getGitRemote(path string) string {

	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()

	return cleanRemoteToOne(string(cmdResponse))
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
