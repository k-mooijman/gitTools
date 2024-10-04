package lib

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher
var err error

func FileWatcher() {
	// Create new watcher.

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
				if event.Has(fsnotify.Rename) {
					log.Println("renamed file:", event.Name)
				}
				if event.Has(fsnotify.Remove) {
					log.Println("deleted file:", event.Name)
				}
				if event.Has(fsnotify.Chmod) {
					log.Println("chmod file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	basePath := "/home/kasper/development/kasper/projects/gitTools"

	//Add a path.
	err = watcher.Add(basePath + "/.git")
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(basePath + "/.git/refs/remotes/origin")
	if err != nil {
		log.Fatal(err)
	}

	//Block main goroutine forever.
	<-make(chan struct{})
}

func Add(path string) {
	if watcher == nil {
		return
	}
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
}
