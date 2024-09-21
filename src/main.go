package main

import (
	"fmt"

	"gitTool/src/lib"
)

func main() {

	myRepo := lib.InitRepos()

	fmt.Printf(" ___________________ \n\n")
	gitRepos := lib.GetGitRepos("/home/kasper/", myRepo)
	fmt.Printf("\n\n ___________________ \n\n")

	for _, value := range gitRepos {
		fmt.Printf("%v \n", value)
	}

	myRepo.GetAllInfo()
	fmt.Printf("******************************* \n")

	myRepo.List()

	fmt.Printf("******************************* \n")

	//yellow := color.New(color.FgYellow).SprintFunc()
	//red := color.New(color.FgRed).SprintFunc()
	//fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

}
