package models

type Market struct {
	Id           int    `json:"id"`
	StartTime    string `json:"startTime"`
	AuctioneerId int    `json:"auctioneerId"`
	GroupId      int    `json:"groupId"`
	Disabled     bool   `json:"disabled"`
}
