package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wineSong/static"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(static.Cors())
	router.StaticFS("/upload", http.Dir("./upload"))
	userRouter(router)
	viewsRouter(router)
	uploadRouter(router)
	auctionRouter(router)
	groupRouter(router)
	marketRouter(router)
	err := router.Run(":8889")
	log.Println(err)
}
