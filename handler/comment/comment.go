package comment

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
	ListByQueryParam(ctx *gin.Context)
	Query(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}

func NewAPI() API {
	return APIImpl{
		projectDao:        dao.NewProjectDao(db.DB),
		userDao:           dao.NewUserDao(db.DB),
		projectHistoryDao: dao.NewProjectHistoryDao(db.DB),
		commentDao:        dao.NewCommentDao(db.DB),
	}
}

type APIImpl struct {
	projectDao        dao.ProjectDaoIF
	userDao           dao.UserDaoIF
	projectHistoryDao dao.ProjectHistoryDaoIF
	commentDao        dao.CommentDaoIF
}

func (A APIImpl) Create(ctx *gin.Context) {
	var request CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		resp.SendError(ctx, err)
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
	comment := dao.Comment{
		Content:      request.Content,
		ResourceID:   request.ResourceID,
		ResourceType: request.ResourceType,
		CommentTime:  time.Now().Unix(),
		Nickname:     user.NickName,
		Permmit:      0,
		ReplierID:    int(user.ID),
	}
	if err := A.commentDao.Create(&comment); err != nil {
		resp.SendError(ctx, err)
		return
	}
	err = A.userDao.Update("score", user.Score+1, "ID = ?", user.ID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, gin.H{
		"id": comment.ID,
	})
}

func (A APIImpl) ListByQueryParam(ctx *gin.Context) {
	resourceID, err := strconv.Atoi(ctx.Query("resource_id"))
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resourceType := ctx.Query("resource_type")
	items, err := A.commentDao.Query(&dao.Comment{
		ResourceID:   resourceID,
		ResourceType: resourceType,
	})
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	result := make([]QueryListResponseItem, 0)

	for _, item := range items {
		user, err := A.userDao.QueryOne("ID = ?", item.ReplierID)
		if err != nil {
			resp.SendError(ctx, err)
			return
		}
		result = append(result, QueryListResponseItem{
			Avatar:      user.Avatar,
			Nickname:    user.NickName,
			Content:     item.Content,
			CommentTime: time.Unix(item.CommentTime, 0).String(),
		})
	}
	resp.SendData(ctx, result)

}

func (A APIImpl) Query(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) Retrieve(ctx *gin.Context) {
	panic("implement me")
}
