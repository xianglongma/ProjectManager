package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/dao/db"
)

type UserAPI interface {
	UserLogin(ctx *gin.Context)
	UserRegister(ctx *gin.Context)
	UserLogout(ctx *gin.Context)
	CurrentUserInfo(ctx *gin.Context)
	UserList(ctx *gin.Context)
	UserOrderList(ctx *gin.Context)
	UserUpdateInfo(ctx *gin.Context)
}

type UserApiImpl struct {
	userDao dao.UserDaoIF
}

func NewUserAPI() UserAPI {
	return &UserApiImpl{
		userDao: dao.NewUserDao(db.DB),
	}
}
