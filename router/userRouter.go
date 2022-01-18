package router

import (
	"github.com/gin-gonic/gin"
	"wineSong/controller"
	"wineSong/middleware"
)

func userRouter(router *gin.Engine) {
	router.POST("/Login", controller.Login)
	router.POST("/Register", controller.Register)
	r := router.Group("/user").Use(middleware.Validate())
	{
		r.GET("/UserInfo", controller.UserInfo)
		r.GET("/SearchAuctioneer", controller.SearchAuctioneer)
		r.POST("/UpdateUser", controller.UpdateUser)
		r.GET("/SearchAllUser", controller.SearchAllUser)
	}
}
