package dao

import (
	"github.com/xianglongma/ProjectManager/dao/db"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title          string `json:"title"` // 项目标题
	Url            string `json:"url"`
	Icon           string `json:"icon"`
	Start          int    `json:"start"`
	End            int    `json:"end"`
	UserID         int    `json:"user_id"` // 项目拥有者
	NickName       string `json:"nickname"`
	Description    string `json:"description"`
	Detail         string `json:"detail"`
	CurrentProcess string `json:"current_process"`  // 当前项目进度
	AllProcessNode string `json:"all_process_node"` // 项目进度节点
	Stars          int    `json:"stars"`            // 项目收藏数
	Scores         int    `json:"scores"`
	CommentID      int    `json:"comment_id"` // 关联评论id
	HistoryID      int    `json:"history_id"` // 历史评论id
	Permmit        int    `json:"permmit"`
	Users          string `json:"users"`
}

func (Project) TableName() string {
	return "projects"
}

var ProjectDao ProjectDaoIF

//go:generate mockgen -source project.go  --destination /mocks/project_mock.go --package dao
type ProjectDaoIF interface {
	AutoMigrate()
	Create(project *Project) error
	Update(field string, value interface{}, where string, args ...interface{}) error
	Updates(project *Project, where string, args ...interface{}) error
	Delete(project *Project) error
	Query(project *Project, limit int, offset int, orderType string, where string, args ...interface{}) ([]Project, error)
	QueryOne(where string, args ...interface{}) (Project, error)
}

func NewProjectDao(db db.Client) ProjectDaoIF {
	ProjectDao = projectDao{client: db}
	// 自动建表
	ProjectDao.AutoMigrate()
	return ProjectDao
}

type projectDao struct {
	client db.Client
}

func (p projectDao) AutoMigrate() {
	p.client.DB().AutoMigrate(&Project{})
}

func (p projectDao) Create(project *Project) error {
	d := p.client.DB().Create(project)
	return d.Error
}

func (p projectDao) Update(field string, value interface{}, where string, args ...interface{}) error {
	result := p.client.DB().Model(&Project{}).Where(where, args).Update(field, value)
	return result.Error
}

func (p projectDao) Updates(project *Project, where string, args ...interface{}) error {
	result := p.client.DB().Model(&Project{}).Where(where, args).Updates(project)
	return result.Error
}

func (p projectDao) Delete(project *Project) error {
	panic("project update")
}

func (p projectDao) Query(project *Project, limit int, offset int, orderType string, where string, args ...interface{}) ([]Project, error) {
	var projects []Project
	db := p.client.DB()
	switch orderType {
	case "time":
		db = db.Order("created_at desc")
	case "score":
		db = db.Order("scores desc")
	case "star":
		db = db.Order("stars desc")
	default:
		db = db.Order("created_at desc")
	}
	db = db.Limit(limit).Offset(offset).Where(project)
	if where != "" {
		db = db.Where(where, args)
	}
	db.Find(&projects)
	return projects, db.Error
}

func (p projectDao) QueryOne(where string, args ...interface{}) (Project, error) {
	var project Project
	result := p.client.DB().First(&project, where, args)
	return project, result.Error
}
