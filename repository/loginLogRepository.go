package repository

import (
	"wineSong/connect"
	"wineSong/models"
)

func CreateLoginLog(log models.LoginLog) error {
	return connect.Db.Create(&log).Error
}
