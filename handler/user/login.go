package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/common"
	"github.com/xianglongma/ProjectManager/pkg/resp"
	"log"
)

func (u *UserApiImpl) UserLogin(ctx *gin.Context) {
	var request UserLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		resp.SendError(ctx, err)
		return
	}
	user, err := u.userDao.QueryOne("nick_name = ?", request.NickName)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	if user.ID == 0 {
		resp.SendError(ctx, resp.InvalidParam)
		return
	}
	if request.Password != user.Password {
		resp.SendError(ctx, resp.InvalidPassword)
		return
	}
	token, expiredAt, err := common.ReleaseToken(&user)
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
