package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gitTool/pkg/git"
	"github.com/gorilla/mux"
)

type server struct {
	repos *git.Repos
}
type action struct {
	Action string `json:"action"`
}

func (s *server) getRepositories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.repos.Repos)
}

func (s *server) getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if id == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.repos.Repos)
		return
	}

	// find the book with the given id
	for _, book := range s.repos.Repos {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "repositorie not found", http.StatusNotFound)
}

// Add a new book
func (s *server) executeAction(w http.ResponseWriter, r *http.Request) {
	var action action
	_ = json.NewDecoder(r.Body).Decode(&action)

	fmt.Printf("Body = %v \n", r.Body)
	fmt.Printf("Action = %v \n", action.Action)

	switch action.Action {
	case "scan":
		go s.repos.ScanForFolders()
	case "get-info":
		go s.repos.GetAllInfo()
	default:
		fmt.Printf("%s.\n", action.Action)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)
}

func New(ctx context.Context, gitRepos *git.Repos) {
	server := &server{
		repos: gitRepos,
	}

	//// Add some dummy data to start with
	//repos = append(repos, repositorie{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})
	//repos = append(repos, repositorie{ID: 2, Title: "Learning Go", Author: "Jon Bodner"})

	// Initialize the router
	r := mux.NewRouter()

	// Define the endpoints
	r.HandleFunc("/repos/", server.getRepositories).Methods("GET")
	r.HandleFunc("/repos/{id}", server.getBook).Methods("GET")
	r.HandleFunc("/set", server.executeAction).Methods("POST")
	//r.HandleFunc("/repos", executeAction).Methods("POST")

	// Start the server
	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
