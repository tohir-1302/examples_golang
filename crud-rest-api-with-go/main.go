package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-golang/helper"
)

const PORT = ":8088"

func main() {
	fmt.Println("Start server ...")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "Hello, World!")
	})

	server := http.Server{
		Addr:    PORT,
		Handler: routes,
	}
	fmt.Println("http://localhost" + PORT + "/")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
