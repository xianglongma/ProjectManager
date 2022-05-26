package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/pkg/resp"
)

func (u *UserApiImpl) CurrentUserInfo(ctx *gin.Context) {
	// 从上下稳重
	userID, exists := ctx.Get("user_id")
	if !exists || userID == 0 {
		resp.SendError(ctx, resp.Unauthorized)
		return
	}
	user, err := u.userDao.QueryOne("id = ?", userID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, user)
}

func (u *UserApiImpl) UserList(ctx *gin.Context) {
	username := ctx.Query("nickname")
	if username == "" {
		resp.SendError(ctx, resp.Unauthorized)
		return
	}
	wildcard := username + "%"
	users, err := u.userDao.Query("nick_name like ?", wildcard)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, users)
}

func (u *UserApiImpl) UserOrderList(ctx *gin.Context) {
	users, err := u.userDao.QueryOrder(&dao.User{}, 10, 0, "score", "")
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, users)
}

func (u *UserApiImpl) UserUpdateInfo(ctx *gin.Context) {
	var request UserUpdateInfoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		resp.SendError(ctx, err)
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists || userID == 0 {
		resp.SendError(ctx, resp.Unauthorized)
		return
	}
	user := dao.User{
		NickName:    request.NickName,
		Description: request.Description,
		Avatar:      request.Avatar,
		Email:       request.Email,
		Mobile:      request.Mobile,
	}
	if err := u.userDao.Updates(&user, "ID = ?", userID); err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendSuccess(ctx)
}
