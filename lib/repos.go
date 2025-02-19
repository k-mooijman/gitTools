package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var hiMagenta = color.New(color.FgHiMagenta).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

type Repos struct {
	Repos map[string]*Repo
}

type Repo struct {
	Location            string
	DoMonitor           bool
	Remote              string
	LastFetchTime       time.Time
	Branches            []Branch
	Test                string
	CurrentBranch       string
	CurrentBranchAhead  int
	CurrentBranchBehind int
}

type Branch struct {
	Branch string
	Ahead  int
	Behind int
}

func (repos *Repos) addByPath(path string) {

	if _, ok := repos.Repos[path]; !ok {
		var repo Repo
		repo.Location = path
		repos.Repos[path] = &repo
	}

}

func InitRepos() *Repos {
	repos := &Repos{}

	content, err := os.ReadFile("repos.json")
	if err != nil {
		fmt.Printf("No File Creating new")
	} else {
		err = json.Unmarshal(content, &repos)
		if err != nil {
			fmt.Printf("No Content Creating new")
		}
	}

	if repos.Repos == nil {
		repos.Repos = make(map[string]*Repo)
	}

	return repos
}

func (repos *Repos) Store() {
	file, _ := json.MarshalIndent(repos, "", " ")

	//fmt.Printf("  file: %s  \n", file)

	_ = os.WriteFile("repos.json", file, 0644)
	fmt.Printf("Repos written  \n")
}

func (repos *Repos) GetAllInfo() {
	repos.GetBranches()
	//repos.GetRemotes()
	repos.GetCurrentBranch()
	//repos.GetFetchDates()
}

func (repos *Repos) GetBranches() {
	for _, repo := range repos.Repos {
		if !repo.DoMonitor {
			continue
		}
		repo.Branches = getGitGetBranches(repo.Location)
		for _, branch := range repo.Branches {
			behind, ahead, err := getGitCurrentBranchBA(repo.Location, branch.Branch)
			if err != nil {
				fmt.Printf("Failed to get BA info: %v", err)
			}
			branch.Ahead = ahead
			branch.Behind = behind
		}

	}
}

func (repos *Repos) GetRemotes() {
	for _, repo := range repos.Repos {
		if !repo.DoMonitor {
			continue
		}
		repo.Remote = getGitRemote(repo.Location)
	}
}

func (repos *Repos) GetCurrentBranch() {
	for _, repo := range repos.Repos {
		//if !repo.DoMonitor {
		//	continue
		//}
		repo.CurrentBranch = getGitCurrentBranch(repo.Location)
		behind, ahead, err := getGitCurrentBranchBA(repo.Location, repo.CurrentBranch)
		if err != nil {
			fmt.Printf("Failed to get BA info: %v", err)
		}
		repo.CurrentBranchAhead = ahead
		repo.CurrentBranchBehind = behind
	}
}
func (repos *Repos) GetFetchDates() {
	for _, repo := range repos.Repos {
		if !repo.DoMonitor {
			continue
		}
		repo.LastFetchTime = getGitFetchDate(repo.Location)
	}
}
