package git

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getGitGetBranches(path string) []*Branch {
	var branches []*Branch
	cmd := exec.Command("git", "branch")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)

	response, _ = strings.CutSuffix(response, "\n")
	temp := strings.Split(response, "\n")
	for _, branch := range temp {
		val := strings.Trim(branch, " *")
		branches = append(branches, &Branch{Branch: val})
	}

	return branches
}

func getGitCurrentBranch(path string) string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response
}

func getGitCurrentBranchBA(path string, branch string, defaultBranch string) (int, int, error) {
	if branch == defaultBranch || branch == "" || defaultBranch == "" {
		return 0, 0, nil
	}
	// git rev-list --left-right --count origin/master...master
	formattedString := fmt.Sprintf("origin/%s...%s", defaultBranch, branch)

	cmd := exec.Command("git", "rev-list", "--left-right", "--count", formattedString)
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")
	parts := strings.Split(response, "\t")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid parts : %s \n", cmdResponse)
	}

	// Parse each part into an integer
	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
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
		// fmt.Println(err)
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

func getGitStatus(path string) string {
	// git status --porcelain
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	response := string(cmdResponse)
	response, _ = strings.CutSuffix(response, "\n")

	return response
}

func getDefaultBranch(path string) string {
	// git remote show origin | sed -n '/HEAD branch/s/.*: //p'
	cmd := exec.Command("git", "remote", "show", "origin")

	cmd.Dir = path
	cmdResponse, _ := cmd.Output()
	s := string(cmdResponse)

	re := regexp.MustCompile(`(?m)^.*HEAD branch.*:\s*(.*)$`)
	results := re.FindAllStringSubmatch(s, -1)
	if len(results) < 1 {
		return ""
	}
	if len(results[0]) < 2 {
		return ""
	}
	//fmt.Printf("Default branch %s\n", results[0][1])

	return results[0][1]
}
