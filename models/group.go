package models

type Group struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AuctioneerId int    `json:"auctioneerId"`
}
