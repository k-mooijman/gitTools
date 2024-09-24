package lib

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
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

func getGitFetchDate(path string) time.Time {

	cmd := exec.Command("stat", "-c", "%y", ".git/FETCH_HEAD")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	//red := color.New(color.FgRed).SprintFunc()
	//
	//fmt.Printf("  LastFetch: %s \n", red(response))

	response, _ = strings.CutSuffix(response, "\n")

	fetchTime, err := time.Parse("2006-01-02 15:04:05.999999 Z0700", response)
	if err != nil {
		fmt.Println(err)
		return time.Time{}
	}

	return fetchTime
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
