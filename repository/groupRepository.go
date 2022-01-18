package repository

import (
	"wineSong/connect"
	"wineSong/models"
)

func CreateGroup(group models.Group) error {
	return connect.Db.Create(&group).Error
}

func SearchGroupByAuctioneerId(auctioneerId int) []models.Group {
	var group []models.Group
	connect.Db.Where("auctioneer_id = ?", auctioneerId).Find(&group)
	return group
}

func SearchGroup() []models.Group {
	var group []models.Group
	connect.Db.Find(&group)
	return group
}

func SearchGroupById(id int) models.Group {
	var group models.Group
	connect.Db.Where("id = ?", id).Find(&group)
	return group
}
