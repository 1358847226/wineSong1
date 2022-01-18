package controller

import (
	"github.com/gin-gonic/gin"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/service"
	"wineSong/util"
)

func CreateMarket(c *gin.Context) {
	var res entity.Result
	var market models.Market
	if bindErr := c.BindJSON(&market); bindErr != nil {
		res = util.ErrReturn(605, bindErr, "前端传值绑定失败")
	} else {
		userId, ok := c.Get("userId")
		if ok {
			market.AuctioneerId = userId.(int)
			res = service.CreateMarket(market)
		} else {
			res = util.ErrReturn(605, nil, "获取用户失败")
		}
	}
	c.JSON(res.Code, res)
}

func SearchMarket(c *gin.Context) {
	var res entity.Result
	groupId := c.Query("groupId")
	res = service.SearchMarket(groupId)
	c.JSON(res.Code, res)
}

func SearchAllMarket(c *gin.Context) {
	var res entity.Result
	res = service.SearchAllMarket()
	c.JSON(res.Code, res)
}
