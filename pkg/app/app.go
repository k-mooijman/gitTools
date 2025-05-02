package app

import (
	"context"
	"log"

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
	app.Repo = git.InitRepos()
	app.ApiServer = api.New(app.Repo)
	return app
}

func (a *App) MustRegisterWith(registerer prometheus.Registerer) {
	a.ApiServer.MustRegisterWith(prometheus.WrapRegistererWithPrefix("git_tool_", registerer))
}

func (a *App) Run(ctx context.Context) {

	go a.ApiServer.Run(ctx)
	log.Printf("Start Scanning folders \n")
	a.Repo.ScanForFolders()
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
