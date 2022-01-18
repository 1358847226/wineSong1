package models

import (
	"wineSong/util"
)

type Auction struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Images       string `json:"images"`
	StartPrice   string `json:"startPrice"`
	FinalPrice   string `json:"finalPrice"`
	CreateTime   string `json:"createTime"`
	State        string `json:"state"`
	AuctioneerId int    `json:"auctioneerId"`
	MarketId     int    `json:"marketId"`
	Buyer        string `json:"buyer"`
}

type Auction1 struct {
	Id           int      `json:"id"`
	UserId       int      `json:"userId"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Images       []string `json:"images"`
	StartPrice   string   `json:"startPrice"`
	FinalPrice   string   `json:"finalPrice"`
	CreateTime   string   `json:"createTime"`
	State        string   `json:"state"`
	AuctioneerId int      `json:"auctioneerId"`
	MarketId     int      `json:"marketId"`
	Buyer        string   `json:"buyer"`
}

func Auction12Auction(auction1 Auction1) Auction {
	auction := Auction{
		Id:           auction1.Id,
		UserId:       auction1.UserId,
		Title:        auction1.Title,
		Content:      auction1.Content,
		Images:       util.A2S(auction1.Images),
		StartPrice:   auction1.StartPrice,
		FinalPrice:   auction1.FinalPrice,
		CreateTime:   auction1.CreateTime,
		State:        auction1.State,
		AuctioneerId: auction1.AuctioneerId,
		MarketId:     auction1.MarketId,
		Buyer:        auction1.Buyer,
	}
	return auction
}

func Auctiion2Auction1(auction Auction) Auction1 {
	auction1 := Auction1{
		Id:           auction.Id,
		UserId:       auction.UserId,
		Title:        auction.Title,
		Content:      auction.Content,
		Images:       util.S2a(auction.Images),
		StartPrice:   auction.StartPrice,
		FinalPrice:   auction.FinalPrice,
		CreateTime:   auction.CreateTime,
		State:        auction.State,
		AuctioneerId: auction.AuctioneerId,
		MarketId:     auction.MarketId,
		Buyer:        auction.Buyer,
	}
	return auction1
}
