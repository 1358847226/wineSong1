package router

import (
	"github.com/gin-gonic/gin"
	"wineSong/controller"
	"wineSong/middleware"
)

func marketRouter(router *gin.Engine) {
	r := router.Group("/v1").Use(middleware.Validate())
	{
		r.POST("/CreateMarket", controller.CreateMarket)
		r.GET("/SearchMarket", controller.SearchMarket)
		r.GET("/SearchAllMarket", controller.SearchAllMarket)
	}
}
