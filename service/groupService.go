package service

import (
	"strconv"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/repository"
	"wineSong/util"
)

func CreateGroup(group models.Group) entity.Result {
	var res entity.Result
	if err := repository.CreateGroup(group); err != nil {
		res = util.ErrReturn(603, err, "sql错误")
	} else {
		res.Code = 200
		res.Message = "成功"
	}
	return res
}

func SearchGroup(auctioneerId int) entity.Result {
	var res entity.Result
	group := repository.SearchGroupByAuctioneerId(auctioneerId)
	res.Code = 200
	res.Data = group
	return res
}

func SearchAllGroup() entity.Result {
	var res entity.Result
	group := repository.SearchGroup()
	res.Code = 200
	res.Data = group
	return res
}

func SearchGroupById(groupId string) entity.Result {
	var res entity.Result
	id, err := strconv.Atoi(groupId)
	if err == nil {
		res.Code = 200
		res.Data = repository.SearchGroupById(id)
	} else {
		res = util.ErrReturn(603, err, "前端传值绑定失败")
	}
	return res
}
