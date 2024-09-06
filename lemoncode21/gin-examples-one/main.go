package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tohir-1302-examples/user"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/:name/:phone", user.GetUserInfo)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
