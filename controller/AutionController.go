package controller

import (
	"github.com/gin-gonic/gin"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/service"
	"wineSong/util"
)

func CreateAuction(c *gin.Context) {
	var auction1 models.Auction1
	var res entity.Result
	if bindErr := c.BindJSON(&auction1); bindErr != nil {
		res = util.ErrReturn(605, bindErr, "前端传值绑定失败")
	} else {
		userId, ok := c.Get("userId")
		if ok {
			auction1.UserId = userId.(int)
			res = service.CreateAuction(models.Auction12Auction(auction1))
		} else {
			res = util.ErrReturn(401, nil, "用户信息错误")
		}
	}
	c.JSON(res.Code, res)
}

func SearchAuctionSend(c *gin.Context) {
	userId, ok := c.Get("userId")
	var res entity.Result
	if ok {
		res = service.SearchAuctionSend(userId.(int))
	} else {
		res = util.ErrReturn(601, nil, "用户信息错误")
	}
	c.JSON(res.Code, res)
}

func SearchAuctionGet(c *gin.Context) {
	userId, ok := c.Get("userId")
	condition := c.Query("condition")
	sort := c.Query("sort")
	var res entity.Result
	if ok {
		res = service.SearchAuctionGet(userId.(int), condition, sort)
	} else {
		res = util.ErrReturn(601, nil, "用户信息错误")
	}
	c.JSON(res.Code, res)
}

func DeleteAuction(c *gin.Context) {
	var res entity.Result
	userId, ok := c.Get("userId")
	if ok {
		auctionId := c.Query("auctionId")
		res = service.DeleteAuction(userId.(int), auctionId)
	} else {
		res = util.ErrReturn(605, nil, "用户信息错误")
	}
	c.JSON(res.Code, res)
}

func SearchAuctionById(c *gin.Context) {
	var res entity.Result
	auctionId := c.Query("auctionId")
	userId, ok := c.Get("userId")
	if ok {
		res = service.SearchAuctionById(userId.(int), auctionId)
	} else {
		res = util.ErrReturn(605, nil, "用户信息错误")
	}
	c.JSON(res.Code, res)
}

func SearchAuctionByMarket(c *gin.Context) {
	var res entity.Result
	marketId := c.Query("marketId")
	res = service.SearchAuctionByMarket(marketId)
	c.JSON(res.Code, res)
}
