package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/common"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/pkg/resp"
	"log"
)

func (u *UserApiImpl) UserRegister(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		resp.SendError(ctx, err)
		return
	}
	user := &dao.User{
		NickName:    request.NickName,
		Mobile:      request.Mobile,
		Password:    request.Password,
		Description: request.Description,
	}
	if err := u.userDao.Create(user); err != nil {
		resp.SendError(ctx, err)
	}
	token, expiredAt, err := common.ReleaseToken(user)
	if err != nil {
		resp.SendError(ctx, err)
		log.Printf("generate token err is %v", err)
		return
	}
	resp.SendData(ctx, RegisterResponse{
		Token:     token,
		ExpiredAt: expiredAt,
	})
}
