package article

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/dao"
	"github.com/xianglongma/ProjectManager/dao/db"
	"github.com/xianglongma/ProjectManager/pkg/resp"
)

type API interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	QueryArtile(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}

func NewAPI() API {
	return APIImpl{
		artileDao: dao.NewArticleDao(db.DB),
		userDao:   dao.NewUserDao(db.DB),
	}
}

type APIImpl struct {
	userDao   dao.UserDaoIF
	artileDao dao.ArticleDaoIF
}

func (A APIImpl) Create(ctx *gin.Context) {
	var request CreateRequest
	if err := ctx.BindJSON(&request); err != nil || request.Content == "" {
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
	article := dao.Article{
		Title:    request.Title,
		Content:  request.Content,
		UserID:   int(user.ID),
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		Stars:    0,
	}
	if err := A.artileDao.Create(&article); err != nil {
		resp.SendError(ctx, err)
		return
	}
	err = A.userDao.Update("score", user.Score+2, "ID = ?", user.ID)
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, gin.H{
		"id": article.ID,
	})
}

func (A APIImpl) List(ctx *gin.Context) {
	orderType := ctx.Query("type")
	owner := ctx.Query("owner")
	projects, err := A.artileDao.Query(&dao.Article{Nickname: owner}, 20, 0, orderType, "")
	if err != nil {
		resp.SendError(ctx, err)
		return
	}
	resp.SendData(ctx, projects)
}

func (A APIImpl) QueryArtile(ctx *gin.Context) {
	panic("implement me")
}

func (A APIImpl) Retrieve(ctx *gin.Context) {
	panic("implement me")
}
