package lib

import (
	"fmt"

	"github.com/fatih/color"
)

type Repos struct {
	Repos map[string]*Repo
}

type Repo struct {
	Location      string
	Remote        string
	CurrentBranch string
	Branches      []string
	Test          string
}

type Branch struct {
	Branch string
}

func (repos *Repos) addByPath(path string) {
	var repo Repo
	repo.Location = path

	repos.Repos[path] = &repo
}

func (repos *Repos) List() {
	for _, repo := range repos.Repos {
		yellow := color.New(color.FgYellow).SprintFunc()
		//red := color.New(color.FgRed).SprintFunc()
		Magenta := color.New(color.FgHiMagenta).SprintFunc()
		//green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("Repo: %s \n", yellow(repo.Location))
		fmt.Printf("  Remote: %s \n", yellow(repo.Remote))
		fmt.Printf("  Current: %s \n", Magenta(repo.CurrentBranch))
		//for _, branch := range repo.Branches {
		//	fmt.Printf("    Branch %s \n", green(branch))
		//}
	}
}

func InitRepos() *Repos {
	repos := &Repos{}
	repos.Repos = make(map[string]*Repo)
	return repos
}

func (repos *Repos) GetAllInfo() {
	repos.GetBranches()
	repos.GetRemotes()
	repos.GetCurrentBranch()
}

func (repos *Repos) GetBranches() {
	for _, repo := range repos.Repos {
		repo.Branches = getGitGetBranches(repo.Location)
	}
}

func (repos *Repos) GetRemotes() {
	for _, repo := range repos.Repos {
		repo.Remote = getGitRemote(repo.Location)
	}
}

func (repos *Repos) GetCurrentBranch() {
	for _, repo := range repos.Repos {
		repo.CurrentBranch = getGitCurrentBranch(repo.Location)
	}
}
