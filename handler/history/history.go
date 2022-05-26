package history

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/dao/db"
	"github.com/xianglongma/ProjectManager/pkg/resp"
	"strconv"
	"time"
)

type API interface {
	Create(ctx *gin.Context)
	ListByID(ctx *gin.Context)
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
	var request CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		resp.SendError(ctx, resp.InvalidParam)
		return
	}
	userID, existed := ctx.Get("user_id")
	if !existed {
		resp.SendError(ctx, resp.Unauthorized)
		return
	}
	user, err := A.userDao.QueryOne("ID = ?", userID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}

	history := dao.ProjectHistory{
		FileUrl:        request.Url,
		Description:    request.Description,
		ProjectID:      request.ProjectID,
		ProcessNode:    request.CurrentProcess,
		UserID:         int(user.ID),
		ModifyUserName: user.NickName,
		ModifyTime:     time.Now().Unix(),
	}
	err = A.projectHistoryDao.Create(&history)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	err = A.projectDao.Updates(&dao.Project{
		Url:            request.Url,
		CurrentProcess: request.CurrentProcess,
	}, "ID = ?", request.ProjectID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendSuccess(ctx)
}

func (A APIImpl) ListByID(ctx *gin.Context) {
	projectID, err := strconv.Atoi(ctx.Query("project_id"))
	if err != nil || projectID < 1 {
		resp.SendError(ctx, resp.Unauthorized)
		return
	}
	items, err := A.projectHistoryDao.Query(10, 0, "project_id = ?", projectID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, items)
}

func (A APIImpl) Query(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) Retrieve(ctx *gin.Context) {
	panic("implement me")
}
