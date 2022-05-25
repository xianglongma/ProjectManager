package dao

import (
	"github.com/xianglongma/ProjectManager/dao/db"
	"gorm.io/gorm"
)

type ProjectHistory struct {
	gorm.Model
	FileUrl        string `json:"file_url"`
	UserID         int    `json:"user_id"`
	ModifyUserName int    `json:"modify_user_name"`
	Description    string `json:"description"`
	ProjectID      int    `json:"project_id"`
	ProcessNode    string `json:"process_node"`
	ModifyTime     int    `json:"modify_time"`
}

func (ProjectHistory) TableName() string {
	return "projects"
}

var ProjectHistoryDao ProjectHistoryDaoIF

type ProjectHistoryDaoIF interface {
	AutoMigrate()
	Create(project *ProjectHistory) error
	Update(project *ProjectHistory) error
	Delete(project *ProjectHistory) error
	Query(where string, args ...interface{}) ([]ProjectHistory, error)
	QueryOne(where string, args ...interface{}) (ProjectHistory, error)
}

func NewProjectHistoryDao(db db.Client) ProjectHistoryDaoIF {
	ProjectHistoryDao = projectHistoryDao{client: db}
	// 自动建表
	ProjectHistoryDao.AutoMigrate()
	return ProjectHistoryDao
}

type projectHistoryDao struct {
	client db.Client
}

func (p projectHistoryDao) AutoMigrate() {
	p.client.DB().AutoMigrate(&Project{})
}

func (p projectHistoryDao) Create(project *ProjectHistory) error {
	d := p.client.DB().Create(project)
	return d.Error
}

func (p projectHistoryDao) Update(project *ProjectHistory) error {
	panic("implement me")
}

func (p projectHistoryDao) Delete(project *ProjectHistory) error {
	panic("implement me")
}

func (p projectHistoryDao) Query(where string, args ...interface{}) ([]ProjectHistory, error) {
	panic("implement me")
}

func (p projectHistoryDao) QueryOne(where string, args ...interface{}) (ProjectHistory, error) {
	panic("implement me")
}
