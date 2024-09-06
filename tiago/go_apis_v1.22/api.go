package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
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
	router.HandleFunc("GET /user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		_, err := w.Write([]byte("User ID: " + userID))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
	})

	middlewareChain := MiddlewareChain(
		RequireAuthMiddleware,
		RequestLoggerMiddleware,
	)

	server := http.Server{
		Addr:    a.addr,
		Handler: middlewareChain(router),
	}

	fmt.Printf("Serving API on %s", a.addr)
	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}
