package main

import (
	_ "embed"

	"gitTool/pkg/app"
)

//go:generate sh -c "echo '0.1' > version.txt"
//go:embed version.txt
var version string

func main() {

	gtApp := app.New("test")
	gtApp.Run()

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
