package main

import (
	_ "gin_frame/app"
	"gin_frame/routes"
)

func main() {
	gin := routes.IniterRouter()
	gin.Run() // gin.Run(":8100") // 监听并在 0.0.0.0:8100 上启动服务
}
