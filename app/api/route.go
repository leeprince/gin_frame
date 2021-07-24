package api

import (
	v1 "gin_frame/app/api/v1"
	"gin_frame/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	routes.RegisterRouter(Routes)
}

func Routes(g *gin.Engine) {
	hello := g.Group("v1")
	{
		hello.GET("/hello", v1.Hello)
		hello.GET("/getUser", v1.GetUser)
	}
}