package router

import (
	"github.com/gin-gonic/gin"
	"wineSong/controller"
	"wineSong/middleware"
)

func groupRouter(router *gin.Engine) {
	r := router.Group("/v1").Use(middleware.Validate())
	{
		r.GET("/SearchGroup", controller.SearchGroup)
		r.GET("/SearchAllGroup", controller.SearchAllGroup)
		r.GET("/SearchGroupById", controller.SearchGroupById)
		r.POST("/CreateGroup", controller.CreateGroup).Use(middleware.Admin())
	}
}
