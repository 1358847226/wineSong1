package repository

import (
	"wineSong/connect"
	"wineSong/models"
)

func CreateAuction(auction models.Auction) error {
	return connect.Db.Create(&auction).Error
}

func SearchAuctionByUserId(userId int) []models.Auction {
	var auction []models.Auction
	connect.Db.Where("user_id = ?", userId).Find(&auction)
	return auction
}

func SearchAuctionByAuctioneerId(auctioneerId int, condition string, sort string) []models.Auction {
	var auction []models.Auction
	connect.Db.Where("auctioneer_id = ? and state like ?", auctioneerId, "%"+condition+"%").Order("market_id " + sort).Find(&auction)
	return auction
}

func SearchAuctionById(id int) models.Auction {
	var auction models.Auction
	connect.Db.Where("id = ?", id).Find(&auction)
	return auction
}

func DeleteAuctionById(id int) error {
	var auction models.Auction
	return connect.Db.Where("id = ?", id).Delete(&auction).Error
}

func SearchAuctionByMarket(id int) []models.Auction {
	var auction []models.Auction
	connect.Db.Where("market_id = ?", id).Find(&auction)
	return auction
}

func UpdateAuction(auction models.Auction) error {
	return connect.Db.Debug().Model(&auction).Save(&auction).Error
}

func SearchAllMarket() []models.Market {
	var market []models.Market
	connect.Db.Find(&market)
	return market
}
