package v1

import (
	"fmt"
	"gin_frame/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context)  {
	fmt.Println("GetUsers - controller")
	userId := c.Query("userId")
	userList := services.GetUsers(map[string]interface{}{
		"id": userId,
	})
	// c.JSON(http.StatusOK, userList)
	c.JSON(http.StatusOK, gin.H{
		"userList": userList,
	})
}