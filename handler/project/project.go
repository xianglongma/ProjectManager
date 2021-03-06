package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/dao/db"
	"github.com/xianglongma/ProjectManager/pkg/resp"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type API interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Query(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}

type APIImpl struct {
	projectDao dao.ProjectDaoIF
	userDao    dao.UserDaoIF
}

func NewAPI() API {
	return APIImpl{
		projectDao: dao.NewProjectDao(db.DB),
		userDao:    dao.NewUserDao(db.DB),
	}
}

func (A APIImpl) Create(ctx *gin.Context) {
	var request CreateRequest
	if err := ctx.BindJSON(&request); err != nil || request.Title == "" {
		resp.SendError(ctx, resp.InvalidParam)
		return
	}
	projects, err := A.projectDao.QueryOne("title = ?", request.Title)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			resp.SendError(ctx, err)
			return
		}
	}
	if projects.ID != 0 {
		resp.SendError(ctx, resp.InvalidProject)
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
	currentNodes := strings.Split(request.AllProcessNode, "|")
	currentNode := currentNodes[1]
	project := dao.Project{
		Title:          request.Title,
		Url:            request.Url,
		Icon:           request.AvatarUrl,
		Description:    request.Description,
		Detail:         request.Detail,
		Permmit:        request.Permmit,
		Users:          request.Users,
		Start:          request.Start,
		End:            request.End,
		CurrentProcess: currentNode,
		AllProcessNode: request.AllProcessNode,
		Stars:          0,
		Scores:         0,
		CommentID:      0,
		HistoryID:      0,
		UserID:         int(userID.(uint)),
		NickName:       user.NickName,
	}
	err = A.projectDao.Create(&project)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	err = A.userDao.Update("score", user.Score+2, "ID = ?", user.ID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, gin.H{
		"id": project.ID,
	})
}

func (A APIImpl) List(ctx *gin.Context) {
	orderType := ctx.Query("type")
	owner := ctx.Query("owner")
	projectNameLike := "%" + ctx.Query("project_name_like") + "%"
	projects, err := A.projectDao.Query(&dao.Project{NickName: owner}, 20, 0, orderType, "title like ?", projectNameLike)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, projects)
}

func (A APIImpl) Query(ctx *gin.Context) {
}

func (A APIImpl) Retrieve(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		resp.SendError(ctx, resp.InvalidParam)
		return
	}
	project, err := A.projectDao.QueryOne("ID = ?", id)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, project)
}
