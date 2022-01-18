package service

import (
	"strconv"
	"time"
	"wineSong/entity"
	"wineSong/models"
	"wineSong/repository"
	"wineSong/util"
)

func CreateMarket(market models.Market) entity.Result {
	var res entity.Result
	if err := repository.CreateMarket(market); err != nil {
		res = util.ErrReturn(603, err, "sql错误")
	} else {
		res.Code = 200
		res.Message = "成功"
	}
	return res
}

func SearchMarket(groupId string) entity.Result {
	var res entity.Result
	if groupId == "" {
		res.Code = 200
		res.Data = repository.SearchMarket()
	} else {
		id, err := strconv.Atoi(groupId)
		if err != nil {
			res = util.ErrReturn(605, err, "系统错误")
		} else {
			startTime := time.Now().Format("2006年01月02日15时")
			res.Code = 200
			market := repository.SearchMarketByGroupIdAndTime(id, startTime)
			for i, v := range market {
				auction := repository.SearchAuctionByMarket(v.Id)
				if len(auction) >= 20 {
					market[i].Disabled = true
				}
			}
			res.Data = market
		}
	}
	return res
}

func SearchAllMarket() entity.Result {
	var res entity.Result
	market := repository.SearchAllMarket()
	res.Code = 200
	res.Data = market
	return res
}
