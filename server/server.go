package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/handler/article"
	"github.com/xianglongma/ProjectManager/handler/comment"
	"github.com/xianglongma/ProjectManager/handler/file"
	"github.com/xianglongma/ProjectManager/handler/history"
	"github.com/xianglongma/ProjectManager/handler/navbar"
	"github.com/xianglongma/ProjectManager/handler/project"
	"github.com/xianglongma/ProjectManager/handler/user"
	"github.com/xianglongma/ProjectManager/middleware"
	"net/http"
)

type Server interface {
	SetRouter() Server
	RunServer()
	Engine() *gin.Engine
}

func NewGinServer(addr string) Server {
	g := gin.New()
	//gin.SetMode(gin.DebugMode)
	return &GinServer{
		engine:     g,
		addr:       addr,
		userAPI:    user.NewUserAPI(),
		navAPI:     navbar.NewNavAPI(),
		fileAPI:    file.NewFileAPI(),
		projectAPI: project.NewAPI(),
		historyAPI: history.NewAPI(),
		commentAPI: comment.NewAPI(),
		articleAPI: article.NewAPI(),
	}
}

type GinServer struct {
	engine     *gin.Engine
	addr       string
	userAPI    user.UserAPI
	navAPI     navbar.NavAPI
	fileAPI    file.API
	projectAPI project.API
	historyAPI history.API
	commentAPI comment.API
	articleAPI article.API
}

func (g *GinServer) SetRouter() Server {
	// 允许跨域请求
	g.engine.Use(middleware.Cors())

	g.engine.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("pong"))
	})

	unAuthGroup := g.engine.Group("/api")
	{
		unAuthGroup.POST("login", g.userAPI.UserLogin)       // /api/login
		unAuthGroup.POST("register", g.userAPI.UserRegister) // /api/register
		unAuthGroup.GET("file/download", g.fileAPI.Download)
		unAuthGroup.GET("image/download", g.fileAPI.DownloadImg)
	}
	// 查看是否登陆
	authGroup := g.engine.Group("/api")
	authGroup.Use(middleware.AuthMiddleWare())
	{
		authGroup.POST("logout", g.userAPI.UserLogout)
		authGroup.POST("user", g.userAPI.UserUpdateInfo)
		authGroup.GET("userinfo", g.userAPI.CurrentUserInfo) // /api/userinfo 获取用户信息
		authGroup.GET("navbar", g.navAPI.GetMenuList)
		authGroup.GET("user/list", g.userAPI.UserList)
		authGroup.GET("user/order/list", g.userAPI.UserOrderList)
	}
	fileGroup := authGroup.Group("/file")
	{
		fileGroup.POST("upload", g.fileAPI.Upload)
	}
	projectGroup := authGroup.Group("/project")
	{
		projectGroup.POST("", g.projectAPI.Create)
		projectGroup.GET("/:id", g.projectAPI.Retrieve)
		projectGroup.GET("/list", g.projectAPI.List)
	}
	historyGroup := authGroup.Group("/history")
	{
		historyGroup.POST("", g.historyAPI.Create)
		historyGroup.GET("list", g.historyAPI.ListByID)
	}
	commentGroup := authGroup.Group("/comment")
	{
		commentGroup.POST("", g.commentAPI.Create)
		commentGroup.GET("/list", g.commentAPI.ListByQueryParam)
	}
	articleGroup := authGroup.Group("/article")
	{
		articleGroup.POST("", g.articleAPI.Create)
		articleGroup.GET("/list", g.articleAPI.List)
	}
	return g
}

func (g *GinServer) RunServer() {
	g.engine.Run(g.addr)
}

func (g *GinServer) Engine() *gin.Engine {
	return g.engine
}
