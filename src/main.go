package main

import (
	"fmt"

	"gitTool/src/lib"
)

func main() {

	fmt.Printf(" ___________________ \n\n")
	gitRepos := lib.GetGitRepos("/")
	fmt.Printf("\n\n ___________________ \n\n")

	for _, value := range gitRepos {
		fmt.Printf("%v \n", value)
	}

}
