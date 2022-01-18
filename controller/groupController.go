package controller

import (
	"github.com/gin-gonic/gin"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/service"
	"wineSong/util"
)

func SearchGroup(c *gin.Context) {
	var res entity.Result
	userId, ok := c.Get("userId")
	if ok {
		res = service.SearchGroup(userId.(int))
	} else {
		res = util.ErrReturn(605, nil, "获取用户失败")
	}
	c.JSON(res.Code, res)
}

func SearchAllGroup(c *gin.Context) {
	var res entity.Result
	res = service.SearchAllGroup()
	c.JSON(res.Code, res)
}

func CreateGroup(c *gin.Context) {
	var res entity.Result
	var group models.Group
	if bindErr := c.BindJSON(&group); bindErr != nil {
		res = util.ErrReturn(605, bindErr, "前端传值绑定失败")
	} else {
		res = service.CreateGroup(group)
	}
	c.JSON(res.Code, res)
}

func SearchGroupById(c *gin.Context) {
	groupId := c.Query("groupId")
	var res entity.Result
	res = service.SearchGroupById(groupId)
	c.JSON(res.Code, res)
}
