package app

import (
	"context"
	"fmt"
	"log"

	"gitTool/pkg/api"
	"gitTool/pkg/git"
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

	myRepo := git.InitRepos()

	go api.New(ctx, myRepo)
	myRepo.ScanForFolders()
	log.Printf("Start Scanning folders \n")
	//folders := file.GetGitRepos("/home/kasper/")
	log.Printf("Finished Scanning folders \n")

	//myRepo.AddByPaths(folders)
	//for _, path := range folders {
	//
	//	log.Printf("Found git at : %q\n", path)
	//
	//}

	//
	//
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
