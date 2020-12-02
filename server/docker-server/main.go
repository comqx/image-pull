package main

import (
	"fmt"
	"image-pull/router"
	"image-pull/tcpServer"
	"os"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func startGinHTTP(httpAddr string ) {
	gin.SetMode(gin.DebugMode)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	ginpprof.Wrap(server)
	router.ConfigApiRouter(&server.RouterGroup)
	fmt.Printf("查看server运行状态： http://%s:8080/debug/pprof\n",httpAddr)
	err := server.Run(fmt.Sprintf("%s:8080",httpAddr))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//global.GVA_LOG = log.Zap() // 初始化zap日志库
	serverAddr := os.Getenv("serverAddr")
	if serverAddr ==""{
		serverAddr = "0.0.0.0"
	}
	httpAddr :=os.Getenv("httpAddr")
	if httpAddr ==""{
		httpAddr = "0.0.0.0"
	}
	listen, err := tcpServer.SerInit(serverAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" Listening and serving tcp on %s:30000\n",serverAddr)
	defer listen.Close()
	startGinHTTP(httpAddr)
}
