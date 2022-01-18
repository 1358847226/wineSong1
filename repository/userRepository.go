package repository

import (
	"wineSong/connect"
	"wineSong/models"
)

func SearchUserByName(userName string) models.User {
	var user models.User
	connect.Db.Where("user_name = ?", userName).Find(&user)
	return user
}

func SearchUserById(id int) models.User {
	var user models.User
	connect.Db.Where("id = ?", id).Find(&user)
	return user
}

func SearchUserByPhone(phone string) models.User {
	var user models.User
	connect.Db.Where("phone = ?", phone).Find(&user)
	return user
}

func SearchUserByQq(qq string) models.User {
	var user models.User
	connect.Db.Where("qq = ?", qq).Find(&user)
	return user
}

func Register(user models.User) error {
	return connect.Db.Create(&user).Error
}

func SearchAuctioneer() []models.User {
	var user []models.User
	connect.Db.Select([]string{"user_name", "id"}).Where("role = ?", "auctioneer").Find(&user)
	return user
}

func SearchAuctioneerByUserId(id int) models.User {
	var user models.User
	connect.Db.Select([]string{"user_name", "id"}).Where("role = ? and id = ?", "auctioneer", id).Find(&user)
	return user
}

func UpdateUser(user models.User) error {
	return connect.Db.Model(&user).Where("id = ? ", user.Id).Update(&user).Error
}

func SearchAllUser() []models.User {
	var user []models.User
	connect.Db.Find(&user)
	return user
}
