package app

import (
	"context"
	"fmt"

	"gitTool/pkg/tools"
)

type Ops struct {
	TestString string `json:"testString"`
}

type App struct {
	Ops Ops `json:"ops"`
}

func New(ops Ops) *App {
	return &App{
		Ops: ops,
	}
}

func (a *App) Run(ctx context.Context) {

	fmt.Printf(tools.Green("Added path %q to watcher\n"), a.Ops.TestString)

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

	tools.WaitForQ(ctx)
}
