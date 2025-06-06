package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitTool/pkg/api"
	"gitTool/pkg/git"
	"gitTool/pkg/tools"
	"github.com/prometheus/client_golang/prometheus"
)

type Ops struct {
	TestString string `json:"testString"`
}

type App struct {
	Ops       Ops `json:"ops"`
	Repo      *git.Repos
	ApiServer *api.Server `json:"apiServer"`
}

func New(ops Ops) *App {
	app := &App{
		Ops: ops,
	}
	app.ApiServer = api.New(app.Repo)
	app.Repo = git.InitRepos()
	return app
}

func (a *App) MustRegisterWith(registerer prometheus.Registerer) {
	a.ApiServer.MustRegisterWith(prometheus.WrapRegistererWithPrefix("git_tool_", registerer))
}

func (a *App) Run(ctx context.Context) {
	go a.ApiServer.Run(ctx)
	time.Sleep(10 * time.Second)
	log.Printf("Start Scanning folders \n")
	message := api.Message{Name: "status", Value: "starting"}
	err := a.ApiServer.Sent(message)
	if err != nil {
		fmt.Printf("Error sent start = %v \n", err)
	}
	a.Repo.ScanForFolders()
	log.Printf("Finished Scanning folders \n")
	message = api.Message{Name: "status", Value: "finishing"}
	err = a.ApiServer.Sent(message)
	if err != nil {
		fmt.Printf("Error = %v \n", err)
	}

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
