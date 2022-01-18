package controller

import (
	"github.com/gin-gonic/gin"
	"wineSong/entity"
	"wineSong/service"
	"wineSong/util"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	println("file", file)
	var res entity.Result
	if err != nil {
		res = util.ErrReturn(605, err, "文件接收失败")
	} else {
		res = service.UploadImg(file, c)
	}
	c.JSON(res.Code, res)
}
