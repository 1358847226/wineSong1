package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
	"wineSong/entity"
	"wineSong/util"
)

func UploadImg(file *multipart.FileHeader, c *gin.Context) entity.Result {
	var res entity.Result
	err := os.Mkdir("./upload", os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	filename := file.Filename
	s := strings.LastIndex(filename, ".")
	fileType := filename[s:]
	name := strconv.Itoa(int(time.Now().UnixNano())) + fileType
	err = c.SaveUploadedFile(file, "./upload/"+name)
	if err != nil {
		res = util.ErrReturn(606, err, "文件写入失败")
	} else {
		res.Code = 200
		res.Message = "成功"
		res.Data = entity.GetConfig().Host + "upload/" + name
	}
	return res
}
