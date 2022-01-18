package controller

import (
	"github.com/gin-gonic/gin"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/service"
	"wineSong/util"
)

func Login(c *gin.Context) {
	var user models.User
	var res entity.Result
	if err := c.BindJSON(&user); err != nil {
		res = util.ErrReturn(605, err, "前端传值绑定失败")
	} else {
		res = service.Login(user)
	}
	c.JSON(res.Code, res)
}

func Register(c *gin.Context) {
	var res entity.Result
	var user models.User
	if bindErr := c.BindJSON(&user); bindErr != nil {
		res = util.ErrReturn(605, bindErr, "前端传值绑定失败")
	} else {
		res = service.Register(user)
	}
	c.JSON(res.Code, res)
}

func UserInfo(c *gin.Context) {
	var res entity.Result
	if userId, err := c.Get("userId"); err {
		res = service.GetUserInfo(userId.(int))
	} else {
		res = util.ErrReturn(606, nil, "系统内部错误")
	}
	c.JSON(res.Code, res)
}

func SearchAuctioneer(c *gin.Context) {
	var res entity.Result
	userId := c.Query("userId")
	res = service.SearchAuctioneer(userId)
	c.JSON(res.Code, res)
}

func UpdateUser(c *gin.Context) {
	var res entity.Result
	var user models.User
	userId, ok := c.Get("userId")
	if ok {
		if bindErr := c.BindJSON(&user); bindErr != nil {
			res = util.ErrReturn(605, bindErr, "前端传值绑定失败")
		} else {
			user.Id = userId.(int)
			res = service.UpdateUser(user)
		}
	} else {
		res = util.ErrReturn(605, nil, "前端传值绑定失败")
	}
	c.JSON(res.Code, res)
}

func SearchAllUser(c *gin.Context) {
	var res entity.Result
	res = service.SearchAllUser()
	c.JSON(res.Code, res)
}
