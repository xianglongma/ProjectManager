package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/common"

	"github.com/xianglongma/ProjectManager/pkg/resp"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get authorization header
		tokenString := ctx.GetHeader("Access-Token")
		// 没有token直接返回空值
		if tokenString == "" {
			resp.SendError(ctx, resp.Unauthorized)
			ctx.Abort()
			return
		}
		token, claims, err := common.ParseToken(tokenString)
		if !token.Valid || err != nil {
			resp.SendError(ctx, resp.Unauthorized)
			ctx.Abort()
			return
		}
		// 获取userID，并将userID写入上下文
		userID := claims.UserID
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
