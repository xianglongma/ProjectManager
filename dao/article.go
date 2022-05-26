package dao

import (
	"github.com/xianglongma/ProjectManager/dao/db"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Stars       int    `json:"stars,omitempty"`
}

func (Article) TableName() string {
	return "articles"
}

var ArticleDao ArticleDaoIF

//go:generate mockgen -source article.go  --destination /mocks/article_mock.go --package dao
type ArticleDaoIF interface {
	AutoMigrate()
	Create(project *Article) error
	Update(field string, value interface{}, where string, args ...interface{}) error
	Updates(project *Article, where string, args ...interface{}) error
	Delete(project *Article) error
	Query(project *Article, limit int, offset int, orderType string, where string, args ...interface{}) ([]Article, error)
	QueryOne(where string, args ...interface{}) (Article, error)
}

func NewArticleDao(db db.Client) ArticleDaoIF {
	ArticleDao = articleDao{client: db}
	// 自动建表
	ArticleDao.AutoMigrate()
	return ArticleDao
}

type articleDao struct {
	client db.Client
}

func (a articleDao) AutoMigrate() {
	a.client.DB().AutoMigrate(&Article{})
}

func (a articleDao) Create(project *Article) error {
	d := a.client.DB().Create(project)
	return d.Error
}

func (a articleDao) Update(field string, value interface{}, where string, args ...interface{}) error {
	panic("implement me")
}

func (a articleDao) Updates(project *Article, where string, args ...interface{}) error {
	panic("implement me")
}

func (a articleDao) Delete(project *Article) error {
	panic("implement me")
}

func (a articleDao) Query(project *Article, limit int, offset int, orderType string, where string, args ...interface{}) ([]Article, error) {
	var projects []Article
	db := a.client.DB()
	switch orderType {
	case "time":
		db = db.Order("created_at desc")
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

func (a articleDao) QueryOne(where string, args ...interface{}) (Article, error) {
	panic("implement me")
}
