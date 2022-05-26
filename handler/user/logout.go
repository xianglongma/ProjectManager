package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/pkg/resp"
)

func (u *UserApiImpl) UserLogout(ctx *gin.Context) {
	resp.SendSuccess(ctx)
}
