package main

import (
	"context"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
	"runtime"
	"strings"

	"gitTool/pkg/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/slxh/go/env"
)

//go:generate sh -c "echo '0.1' > version.txt"
//go:embed version.txt
var version string

var errUnsupportedPlatform = errors.New("unsupported platform")

func main() {
	const (
		application = "gitTool"
	)

	var (
		browser bool
	)
	appOps := app.Ops{}

	flag.StringVar(&appOps.TestString, "api-addr", ":8096", "API listen address")
	flag.BoolVar(&browser, "browser", false, "whether to start a browser")
	fmt.Printf("browser flag set to %t\n", browser)

	// flag.Parse()
	if err := env.ParseWithFlags(); err != nil {
		log.Fatal("Invalid environment variables or CLI flag: %v", "err", err)
	}

	fmt.Printf("Added path %q to watcher\n", appOps.TestString)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if browser {
		fmt.Printf("Opening browser \n")
		err := openbrowser("http://localhost:8000/repos/")
		if err != nil {
			fmt.Errorf("failed to open browser: %w", err)
		}
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	gtApp := app.New(appOps)
	gtApp.MustRegisterWith(prometheus.DefaultRegisterer)
	gauge := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: strings.ToLower(application),
			Name:      "build_info",
			Help:      "A metric containing build information for the application.",
			ConstLabels: prometheus.Labels{
				"version":   strings.TrimSpace(version),
				"goversion": runtime.Version(),
				"goos":      runtime.GOOS,
				"goarch":    runtime.GOARCH,
			},
		},
		func() float64 { return 1 },
	)
	prometheus.DefaultRegisterer.MustRegister(gauge)
	gtApp.Run(ctx)

}

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errUnsupportedPlatform
	}

	if err != nil {
		return fmt.Errorf("could not open browser: %w", err)
	}

	return nil
}

//################################################################################
//package main
//
//import (
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//	"sync"
//
//	"github.com/fsnotify/fsnotify"
//)
//
//var (
//	watcher *fsnotify.Watcher
//	mu      sync.Mutex
//)
//
//func init() {
//	var err error
//	watcher, err = fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatalf("Failed to create watcher: %v", err)
//	}
//}
//
//func addPath(path string) error {
//	mu.Lock()
//	defer mu.Unlock()
//
//	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if info.IsDir() && p != path {
//			return watcher.Add(p)
//		}
//		return nil
//	})
//	if err != nil {
//		return fmt.Errorf("failed to add path %q: %v", path, err)
//	}
//	fmt.Printf("Added path %q to watcher\n", path)
//	return nil
//}
//
//func removePath(path string) error {
//	mu.Lock()
//	defer mu.Unlock()
//
//	return watcher.Remove(path)
//}
//
//func startWatching() {
//	go func() {
//		for {
//			select {
//			case event, ok := <-watcher.Events:
//				if !ok {
//					return
//				}
//				fmt.Printf("Event: %s\n", event)
//			case err, ok := <-watcher.Errors:
//				if !ok {
//					return
//				}
//				fmt.Printf("Error: %s\n", err)
//			}
//		}
//	}()
//}
//
//func main() {
//	startWatching()
//
//	// Simulate receiving new paths to watch from another process
//	newPaths := []string{"/home/kasper/temp"}
//
//	for _, path := range newPaths {
//		if err := addPath(path); err != nil {
//			log.Printf("Failed to add path %q: %v", path, err)
//		}
//	}
//
//	// Keep the program running
//	select {}
//}
