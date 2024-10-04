package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Repos struct {
	Repos map[string]*Repo
}

type Repo struct {
	Location      string
	DoMonitor     bool
	Remote        string
	LastFetchTime time.Time
	CurrentBranch string
	Branches      []string
	Test          string
}

type Branch struct {
	Branch string
}

func (repos *Repos) addByPath(path string) {

	if _, ok := repos.Repos[path]; !ok {
		var repo Repo
		repo.Location = path
		repos.Repos[path] = &repo
	}

}

func (repos *Repos) List() {
	for _, repo := range repos.Repos {
		if !repo.DoMonitor {
			continue
		}
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		hiMagenta := color.New(color.FgHiMagenta).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()

		fmt.Printf("Repo: %s \n", yellow(repo.Location))
		emptyTime := time.Time{}
		t := time.Now()
		t2 := t.AddDate(0, 0, -14)
		if repo.LastFetchTime == emptyTime {
			fmt.Printf("  LastFetch: %s \n", magenta("never fetched"))

		} else {
			if repo.LastFetchTime.After(t2) {
				fmt.Printf("  LastFetch: %s \n", green(repo.LastFetchTime))

			} else {
				fmt.Printf("  LastFetch: %s \n", red(repo.LastFetchTime))

			}
		}

		fmt.Printf("  Remote: %s \n", yellow(repo.Remote))
		fmt.Printf("  Current: %s \n", hiMagenta(repo.CurrentBranch))
		for _, branch := range repo.Branches {
			fmt.Printf("    Branch %s \n", green(branch))
		}
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
	repos.GetRemotes()
	repos.GetCurrentBranch()
	repos.GetFetchDates()
}

func (repos *Repos) GetBranches() {
	for _, repo := range repos.Repos {
		if !repo.DoMonitor {
			continue
		}
		repo.Branches = getGitGetBranches(repo.Location)
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
		if !repo.DoMonitor {
			continue
		}
		repo.CurrentBranch = getGitCurrentBranch(repo.Location)
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
