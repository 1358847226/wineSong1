package repository

import (
	"wineSong/connect"
	"wineSong/models"
)

func CreateMarket(market models.Market) error {
	return connect.Db.Create(&market).Error
}

func SearchMarket() []models.Market {
	var market []models.Market
	connect.Db.Find(&market)
	return market
}

func SearchMarketByGroupIdAndTime(id int, startTime string) []models.Market {
	var market []models.Market
	connect.Db.Debug().Order("start_time").Where("group_id = ? and start_time >= ?", id, startTime).Find(&market)
	return market
}
