package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserClaims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func NewToken(uid int) (string, error) {
	expTime := time.Now().Add(time.Hour * 24).Unix()
	claims := &UserClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := "kasdnjklasn"
	return token.SignedString([]byte(key))
}

func Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Request.Header.Get("Access-Token")
		callFun := func(token *jwt.Token) (interface{}, error) {
			return []byte("kasdnjklasn"), nil
		}
		if token, err := jwt.ParseWithClaims(key, &UserClaims{}, callFun); err == nil {
			if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
				ctx.Set("userId", claims.UserID)
				ctx.Next()
			} else {
				ctx.JSON(401, gin.H{
					"code":    401,
					"message": "请先登录",
				})
				ctx.Abort()
			}
		} else {
			ctx.JSON(401, gin.H{
				"code":    401,
				"message": "请先登录",
			})
			ctx.Abort()
		}
	}
}
