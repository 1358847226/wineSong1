package service

import (
	"strconv"
	"time"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/repository"
	"wineSong/util"
)

func CreateAuction(auction models.Auction) entity.Result {
	var res entity.Result
	auction.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	auctioneer := repository.SearchUserById(auction.AuctioneerId)
	if auctioneer.Id == 0 || auctioneer.Role != "auctioneer" {
		res = util.ErrReturn(604, nil, "拍卖员不存在")
	} else {
		if auction.Id == 0 {
			auction.State = "待拍卖"
			if err := repository.CreateAuction(auction); err != nil {
				res = util.ErrReturn(603, err, "sql错误")
			} else {
				res.Code = 200
				res.Message = "成功"
			}
		} else {
			if auction.State != "已成交" {
				auction.Buyer = ""
				auction.FinalPrice = ""
			}
			if err := repository.UpdateAuction(auction); err != nil {
				res = util.ErrReturn(603, err, "sql错误")
			} else {
				res.Code = 200
				res.Message = "成功"
			}
		}

	}
	return res
}

func SearchAuctionSend(userId int) entity.Result {
	var res entity.Result
	auction := repository.SearchAuctionByUserId(userId)
	var auction1 []models.Auction1
	for i := 0; i < len(auction); i++ {
		a1 := models.Auctiion2Auction1(auction[i])
		auction1 = append(auction1, a1)
	}
	res.Code = 200
	res.Data = auction1
	return res
}

func SearchAuctionGet(userId int, condition string, sort string) entity.Result {
	var res entity.Result
	auction := repository.SearchAuctionByAuctioneerId(userId, condition, sort)
	var auction1 []models.Auction1
	for i := 0; i < len(auction); i++ {
		a1 := models.Auctiion2Auction1(auction[i])
		auction1 = append(auction1, a1)
	}
	res.Code = 200
	res.Data = auction1
	return res
}

func DeleteAuction(userId int, auctionId string) entity.Result {
	var res entity.Result
	id, err := strconv.Atoi(auctionId)
	if err != nil {
		res = util.ErrReturn(605, err, "前端传值绑定失败")
	} else {
		auction := repository.SearchAuctionById(id)
		if auction.UserId == userId || auction.AuctioneerId == userId {
			if err1 := repository.DeleteAuctionById(id); err1 != nil {
				res = util.ErrReturn(603, err1, "sql错误")
			} else {
				res.Code = 200
				res.Message = "成功"
			}
		} else {
			res = util.ErrReturn(601, nil, "没有权限")
		}
	}
	return res
}

func SearchAuctionById(userId int, auctionId string) entity.Result {
	var res entity.Result
	id, err := strconv.Atoi(auctionId)
	if err != nil {
		res = util.ErrReturn(605, err, "前端传值绑定失败")
	} else {
		auction := repository.SearchAuctionById(id)
		if auction.UserId == userId || auction.AuctioneerId == userId {
			res.Code = 200
			res.Data = models.Auctiion2Auction1(auction)
		} else {
			res = util.ErrReturn(601, nil, "没有权限")
		}
	}
	return res
}

func SearchAuctionByMarket(marketId string) entity.Result {
	var res entity.Result
	id, err := strconv.Atoi(marketId)
	if err != nil {
		res = util.ErrReturn(605, err, "前端传值绑定失败")
	} else {
		auction := repository.SearchAuctionByMarket(id)
		var auction1 []models.Auction1
		for i := 0; i < len(auction); i++ {
			a := models.Auctiion2Auction1(auction[i])
			auction1 = append(auction1, a)
		}
		res.Code = 200
		res.Data = auction1
	}
	return res
}
