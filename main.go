package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	svr := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("/", api.rootHandler)
	mux.HandleFunc("/users", api.usersHandler)

	log.Println("Starting server on", api.addr)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
