package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tohir-1302/gin_jwt_token/helper"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	})

	server := &http.Server{
		Addr:    ":8091",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
