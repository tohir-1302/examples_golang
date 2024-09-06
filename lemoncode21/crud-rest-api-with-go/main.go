package main

import (
	"fmt"
	"net/http"
	"rest-api-golang/config"
	"rest-api-golang/controller"
	"rest-api-golang/helper"
	"rest-api-golang/repository"
	"rest-api-golang/routers"
	"rest-api-golang/service"
)

const PORT = ":8089"

func main() {
	fmt.Println("Start server ...")
	db := config.DatabaseConnection()
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookServiceImpl(bookRepository)
	bookController := controller.NewBookController(bookService)

	server := http.Server{
		Addr:    PORT,
		Handler: routers.GetRouters(bookController),
	}
	fmt.Println("http://localhost" + PORT + "/")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
