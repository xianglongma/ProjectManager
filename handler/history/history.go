package history

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/dao/db"
)

type API interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Query(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}

func NewAPI() API {
	return APIImpl{
		projectDao:        dao.NewProjectDao(db.DB),
		userDao:           dao.NewUserDao(db.DB),
		projectHistoryDao: dao.NewProjectHistoryDao(db.DB),
	}
}

type APIImpl struct {
	projectDao        dao.ProjectDaoIF
	userDao           dao.UserDaoIF
	projectHistoryDao dao.ProjectHistoryDaoIF
}

func (A APIImpl) Create(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) List(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) Query(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) Retrieve(ctx *gin.Context) {
	panic("implement me")
}
