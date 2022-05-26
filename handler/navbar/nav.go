package navbar

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/pkg/resp"
)

type NavAPI interface {
	GetMenuList(ctx *gin.Context)
}

type NavAPIImpl struct {
}

func NewNavAPI() NavAPI {
	return NavAPIImpl{}
}

func (n NavAPIImpl) GetMenuList(ctx *gin.Context) {
	menus := []MenuList{
		{
			Name:     "Home",
			ParentID: 0,
			ID:       20001,
			Meta: MenuMeta{
				Icon:  "home",
				Title: "主页",
				Show:  true,
			},
			Component: "Home",
		},
		{
			Name:     "ManagerProjects",
			ParentID: 0,
			ID:       20002,
			Meta: MenuMeta{
				Icon:  "project",
				Title: "项目",
				Show:  true,
			},
			Component: "SearchProjects",
		},
		{

			Name:     "BlankArticles",
			ParentID: 0,
			ID:       20003,
			Meta: MenuMeta{
				Icon:  "global",
				Title: "世界",
				Show:  true,
			},
			Component: "SearchArticles",
		},
		{

			Name:     "navAccount",
			ParentID: 0,
			ID:       20004,
			Meta: MenuMeta{
				Icon:  "user",
				Title: "个人",
				Show:  true,
			},
			Redirect:  "/account/center",
			Component: "RouteView",
		},
		{

			Name:     "ArticleDetail",
			ParentID: 0,
			ID:       20005,
			Meta: MenuMeta{
				Icon:  "global",
				Title: "文章",
				Show:  false,
			},
			Component: "ArticleDetail",
		},
		{
			Name:     "account",
			ParentID: 0,
			ID:       10028,
			Meta: MenuMeta{
				Title: "个人页",
				Icon:  "user",
				Show:  false,
			},
			Redirect:  "/account/center",
			Component: "RouteView",
		},
		{
			Name:     "center",
			ParentID: 10028,
			ID:       10029,
			Meta: MenuMeta{
				Title: "个人中心",
				Show:  false,
			},
			Component: "AccountCenter",
		},
		{
			Name:     "settings",
			ParentID: 10028,
			ID:       10030,
			Meta: MenuMeta{
				Title:        "个人设置",
				HideHeader:   true,
				HideChildren: true,
				Show:         false,
			},
			Redirect:  "/account/settings/basic",
			Component: "AccountSettings",
		},
		{
			Name:     "BasicSettings",
			Path:     "/account/settings/basic",
			ParentID: 10030,
			ID:       10031,
			Meta: MenuMeta{
				Title: "基本设置",
				Show:  false,
			},
			Component: "BasicSetting",
		},
		{

			Name:     "Success",
			ParentID: 0,
			ID:       20005,
			Meta: MenuMeta{
				Icon:  "project",
				Title: "项目",
				Show:  false,
			},
			Component: "ResultSuccess",
		},
		{

			Name:     "CreateProject",
			ParentID: 0,
			ID:       20006,
			Meta: MenuMeta{
				Icon:  "project",
				Title: "项目",
				Show:  false,
			},
			Component: "CreateProject",
		},
	}
	resp.SendData(ctx, menus)
}
