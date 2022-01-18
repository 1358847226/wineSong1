package router

import (
	"github.com/gin-gonic/gin"
)

func viewsRouter(router *gin.Engine) {
	router.LoadHTMLGlob("./tmpl/*")
	r := router.Group("/mahjong")
	{
		r.GET("/ComputeScores")
	}
}
