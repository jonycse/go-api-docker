package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/fsnotify/fsnotify"
)

// watch this data path, notify main if file has changes
const dataPath = "/home/jony/Developer/jony/go-api-docker/data"

func main() {
	r := mux.NewRouter()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	go watchData(watcher)

	// watching data path
	err = watcher.Add(dataPath)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", r)
}

func watchData(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("Modified file: %s", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			panic(err)
		}
	}
}
