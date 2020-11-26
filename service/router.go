package service

import "github.com/gin-gonic/gin"

func Router(api *gin.RouterGroup) {
	api.POST("/sendImage", sendImageInfo)
	api.GET("/getRegisteredList", getRegisteredList)
	//api.GET("/ws", wsClient)  // ws
}
