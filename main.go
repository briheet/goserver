package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from GET /users\n"))
}

func (a *api) postUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from POST /users\n"))
}

func (a *api) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.getUserHandler(w, r)
	case http.MethodPost:
		a.postUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *api) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is flashed for / path\n"))
}

func main() {
	api := &api{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", api.rootHandler)
	mux.HandleFunc("/users", api.usersHandler)

	svr := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	log.Println("Starting server on", api.addr)
	svr.ListenAndServe()
}
