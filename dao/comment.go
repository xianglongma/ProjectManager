package dao

import (
	"github.com/xianglongma/ProjectManager/dao/db"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ParentID     int    `json:"parent_id"`
	ReplierID    int    `json:"replier_id"`
	ResourceID   int    `json:"resource_id"`   // 被评论的资源id
	ResourceType string `json:"resource_type"` // 被评论的资源type
	Content      string `json:"content"`       // 评论内容
	Permmit      int    `json:"permmit"`       // 是否被禁用
	CommentTime  int64  `json:"comment_time"`
	Nickname     string `json:"nickname"`
}

func (Comment) TableName() string {
	return "comments"
}

var CommentDao CommentDaoIF

//go:generate mockgen -source comment.go  --destination /mocks/comment_mock.go --package dao
type CommentDaoIF interface {
	AutoMigrate()
	Create(project *Comment) error
	Update(field string, value interface{}, where string, args ...interface{}) error
	Updates(project *Comment, where string, args ...interface{}) error
	Delete(project *Comment) error
	Query(project *Comment) ([]Comment, error)
	QueryOne(where string, args ...interface{}) (Comment, error)
}

func NewCommentDao(db db.Client) CommentDaoIF {
	CommentDao = commentDao{client: db}
	// 自动建表
	CommentDao.AutoMigrate()
	return CommentDao
}

type commentDao struct {
	client db.Client
}

func (c commentDao) AutoMigrate() {
	c.client.DB().AutoMigrate(&Comment{})
}

func (c commentDao) Create(project *Comment) error {
	d := c.client.DB().Create(project)
	return d.Error
}

func (c commentDao) Update(field string, value interface{}, where string, args ...interface{}) error {
	panic("implement me")
}

func (c commentDao) Updates(project *Comment, where string, args ...interface{}) error {
	panic("implement me")
}

func (c commentDao) Delete(project *Comment) error {
	panic("implement me")
}

func (c commentDao) Query(project *Comment) ([]Comment, error) {
	var projects []Comment
	result := c.client.DB().Order("created_at desc").Where(project).Find(&projects)
	return projects, result.Error
}

func (c commentDao) QueryOne(where string, args ...interface{}) (Comment, error) {
	panic("implement me")
}
