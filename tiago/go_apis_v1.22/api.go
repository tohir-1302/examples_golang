package main

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (a *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("POST /user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		_, err := w.Write([]byte("User ID: " + userID))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
	})

	server := http.Server{
		Addr:    a.addr,
		Handler: router,
	}
	fmt.Printf("Serving API on %s", a.addr)
	return server.ListenAndServe()
}
