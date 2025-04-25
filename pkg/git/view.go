package git

import (
	"fmt"

	"gitTool/pkg/tools"
)

func (repos *Repos) ListAllNeedingAction() {
	fmt.Printf("  count: %s \n", tools.Magenta(len(repos.Repos)))

	for _, repo := range repos.Repos {
		//if !repo.DoMonitor {
		//	continue
		//}

		if repo.CurrentBranchAhead > 0 {

			fmt.Printf("Repo: %s \n", tools.Yellow(repo.Location))
			//emptyTime := time.Time{}
			//t := time.Now()
			//t2 := t.AddDate(0, 0, -14)
			//if repo.LastFetchTime == emptyTime {
			//	fmt.Printf("  LastFetch: %s \n", magenta("never fetched"))
			//
			//} else {
			//	if repo.LastFetchTime.After(t2) {
			//		fmt.Printf("  LastFetch: %s \n", green(repo.LastFetchTime))
			//
			//	} else {
			//		fmt.Printf("  LastFetch: %s \n", red(repo.LastFetchTime))
			//
			//	}
			//}

			// fmt.Printf("  Remote: %s \n", yellow(repo.Remote))
			fmt.Printf("  Current: %s \n", tools.HiMagenta(repo.CurrentBranch))
			fmt.Printf("         : %s Behind %s Ahead \n", tools.Green(repo.CurrentBranchBehind), tools.Red(repo.CurrentBranchAhead))

			//for _, branch := range repo.Branches {
			//	fmt.Printf("    Branch %s \n", green(branch))
			//}

		}
	}
}

func (repos *Repos) List() {
	fmt.Printf("  count: %s \n", tools.Magenta(len(repos.Repos)))

	for _, repo := range repos.Repos {
		fmt.Printf("Repo: %s \n", tools.Yellow(repo.Location))

		//if !repo.DoMonitor {
		//	continue
		//}

		//emptyTime := time.Time{}
		//t := time.Now()
		//t2 := t.AddDate(0, 0, -14)
		//if repo.LastFetchTime == emptyTime {
		//	fmt.Printf("  LastFetch: %s \n", magenta("never fetched"))
		//
		//} else {
		//	if repo.LastFetchTime.After(t2) {
		//		fmt.Printf("  LastFetch: %s \n", green(repo.LastFetchTime))
		//
		//	} else {
		//		fmt.Printf("  LastFetch: %s \n", red(repo.LastFetchTime))
		//
		//	}
		//}

		// fmt.Printf("  Remote: %s \n", yellow(repo.Remote))
		fmt.Printf("  Current: %s \n", tools.HiMagenta(repo.CurrentBranch))
		fmt.Printf("         : %s Behind %s Ahead \n", tools.Green(repo.CurrentBranchBehind), tools.Red(repo.CurrentBranchAhead))

		for _, branch := range repo.Branches {
			fmt.Printf("    Branch %s  - %s Behind %s Ahead \n", tools.Green(branch.Branch, branch.Behind, branch.Ahead))
		}

	}
}
