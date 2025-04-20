package app

import (
	"flag"
	"fmt"
	"log"

	"gitlab.com/slxh/go/env"
)

type App struct {
	Ops string `json:"ops"`
}

func New(ops string) *App {
	return &App{
		Ops: ops,
	}
}

func (a *App) Run() {

	var (
		temp string
	)
	flag.StringVar(&temp, "api-addr", ":8096", "API listen address")

	//flag.Parse()
	if err := env.ParseWithFlags(); err != nil {
		log.Fatal("Invalid environment variables or CLI flag: %v", "err", err)
	}

	fmt.Printf("Added path %q to watcher\n", temp)

	//myRepo := lib2.InitRepos()
	////
	////fmt.Printf(" ___________________ \n\n")
	//lib2.GetGitRepos("/home/kasper/", myRepo)
	//myRepo.GetAllInfo()
	////fmt.Printf("******************************* \n")
	//myRepo.ListAllNeedingAction()
	////fmt.Printf("******************************* \n")
	////myRepo.Store()
	////fmt.Printf("******************************* \n")
	//
	//fmt.Printf(" ___________________ \n\n")
	//
	////pkg.FileWatcher()
	////pkg.Add("/home/kasper")
	////pkg.DoFile()
	////pkg.WaitForQ()
	//fmt.Printf(" ___________________ \n\n")
	//myRepo.List()

}
