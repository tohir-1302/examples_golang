package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Info struct {
	Name  string `uri:"name" binding:"required"`
	Phone string `uri:"phone" binding:"required"`
}

func GetUserInfo(c *gin.Context) {
	var user Info
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
