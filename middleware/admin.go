package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wineSong/entity"
	"wineSong/repository"
	"wineSong/util"
)

func Admin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var res entity.Result
		userId, ok := ctx.Get("userId")
		if ok {
			user := repository.SearchUserById(userId.(int))
			if strings.Contains(user.Role, "admin") {
				ctx.Next()
			} else {
				res = util.ErrReturn(602, nil, "没有权限")
				ctx.JSON(res.Code, res)
				ctx.Abort()
			}
		} else {
			res = util.ErrReturn(604, nil, "用户id不存在")
			ctx.JSON(res.Code, res)
			ctx.Abort()
		}
	}
}
