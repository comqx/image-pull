package main

import (
	"fmt"
	"image-pull/router"
	"image-pull/tcpServer"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func startGinHTTP() {
	gin.SetMode(gin.DebugMode)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	ginpprof.Wrap(server)
	router.ConfigApiRouter(&server.RouterGroup)
	fmt.Println("查看server运行状态： http://0.0.0.0:8080/debug/pprof")
	err := server.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//global.GVA_LOG = log.Zap() // 初始化zap日志库
	listen, err := tcpServer.SerInit("0.0.0.0:30000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" Listening and serving tcp on 0.0.0.0:30000")
	defer listen.Close()
	startGinHTTP()
}
