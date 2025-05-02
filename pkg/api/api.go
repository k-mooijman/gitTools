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
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	repos *git.Repos
	get   prometheus.Counter
}
type action struct {
	Action string `json:"action"`
}

func (s *Server) getRepositories(w http.ResponseWriter, r *http.Request) {
	s.get.Inc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.repos.Repos)
}

func (s *Server) getBook(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) executeAction(w http.ResponseWriter, r *http.Request) {
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

func New(gitRepos *git.Repos) *Server {
	server := &Server{
		repos: gitRepos,
		get: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "api_get_hits_total",
			Help: "Number of API get hits",
		}),
	}
	return server
}
func (s *Server) Run(ctx context.Context) {
	// Initialize the router
	r := mux.NewRouter()

	// Define the endpoints
	r.HandleFunc("/repos/", s.getRepositories).Methods("GET")
	r.HandleFunc("/repos/{id}", s.getBook).Methods("GET")
	r.HandleFunc("/set", s.executeAction).Methods("POST")
	//r.HandleFunc("/repos", executeAction).Methods("POST")

	// Start the Server
	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// MustRegisterWith registers DNS client metrics with the given registerer.
func (s *Server) MustRegisterWith(r prometheus.Registerer) {
	r.MustRegister(s.get)
}
