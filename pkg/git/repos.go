package git

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gitTool/pkg/file"
)

type Repos struct {
	Repos map[string]*Repo
}

type Repo struct {
	ID                  int
	DefaultBranch       string
	Location            string
	DoMonitor           bool
	Remote              string
	LastFetchTime       time.Time
	Branches            []*Branch
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

func (repos *Repos) AddByPath(path string) {
	if _, ok := repos.Repos[path]; !ok {
		var repo Repo
		repo.Location = path
		repos.Repos[path] = &repo
	}
}

func (repos *Repos) AddByPaths(folders []string) {
	for _, path := range folders {
		repos.AddByPath(path)
	}
}

func (repos *Repos) ScanForFolders() {
	folders, _ := file.GetGitRepos("/home/kasper/development/kasper/projects/gitTool/")
	repos.AddByPaths(folders)
}

func (repos *Repos) Store() {
	file, _ := json.MarshalIndent(repos, "", " ")
	_ = os.WriteFile("repos.json", file, 0o644)
	fmt.Printf("Repos written  \n")
}

func (repos *Repos) GetAllInfo() {
	repos.GetDefaultBranches()
	repos.GetCurrentBranch()
	// repos.GetRemotes()
	// repos.GetFetchDates()
	repos.GetBranches()
}

func (repos *Repos) GetBranches() {
	for _, repo := range repos.Repos {
		fmt.Printf(".")
		//if !repo.DoMonitor {
		//	continue
		//}
		repo.Branches = getGitGetBranches(repo.Location)
		for _, branch := range repo.Branches {
			if branch.Branch != repo.DefaultBranch {
				fmt.Printf("Get for %s \n", branch.Branch)
				behind, ahead, err := getGitCurrentBranchBA(repo.Location, branch.Branch, repo.DefaultBranch)
				if err != nil {
					fmt.Printf("Failed to get BA info: %v", err)
				}
				branch.Ahead = ahead
				branch.Behind = behind
			}
		}
	}
	fmt.Printf("! \n")
}
func (repos *Repos) GetDefaultBranches() {
	for _, repo := range repos.Repos {
		fmt.Printf(".")
		//if !repo.DoMonitor {
		//	continue
		//}
		repo.DefaultBranch = getDefaultBranch(repo.Location)
	}
	fmt.Printf("! \n")
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
		behind, ahead, err := getGitCurrentBranchBA(repo.Location, repo.CurrentBranch, repo.DefaultBranch)
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
