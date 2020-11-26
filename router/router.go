package router

import (
	"image-pull/service"

	"github.com/gin-gonic/gin"
)

func ConfigApiRouter(r *gin.RouterGroup) {
	api := r.Group("/api")
	service.Router(api)
}
