package routers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-golang/controller"
	"rest-api-golang/helper"
)

func GetRouters(bookController *controller.BookController) *httprouter.Router {
	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := fmt.Fprintf(w, "Hello, World!")
		helper.PanicIfError(err)
	})

	routes.GET("/api/book", bookController.FindAll)
	routes.GET("/api/book/:id", bookController.FindById)
	routes.POST("/api/book", bookController.Create)
	routes.PATCH("/api/book/:id", bookController.Update)
	routes.DELETE("/api/book/:id", bookController.Delete)

	return routes
}
