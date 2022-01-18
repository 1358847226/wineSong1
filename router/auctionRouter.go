package router

import (
	"github.com/gin-gonic/gin"
	"wineSong/controller"
	"wineSong/middleware"
)

func auctionRouter(router *gin.Engine) {
	r := router.Group("/v1").Use(middleware.Validate())
	{
		r.POST("/CreateAuction", controller.CreateAuction)
		r.GET("/SearchAuctionSend", controller.SearchAuctionSend)
		r.GET("/SearchAuctionGet", controller.SearchAuctionGet)
		r.GET("/DeleteAuction", controller.DeleteAuction)
		r.GET("/SearchAuctionById", controller.SearchAuctionById)
		r.GET("/SearchAuctionByMarket", controller.SearchAuctionByMarket)
	}
}
