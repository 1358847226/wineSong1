package router

import (
	"github.com/gin-gonic/gin"
	"wineSong/controller"
	"wineSong/middleware"
)

func uploadRouter(router *gin.Engine) {
	r := router.Group("/upload").Use(middleware.Validate())
	{
		r.POST("/UploadImg", controller.UploadImg)
	}

}
