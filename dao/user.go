package dao

import (
	"github.com/xianglongma/ProjectManager/dao/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
	NickName    string `json:"nickname"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Tags        string `json:"tags"`
	Score       int    `json:"score"`
}

func (User) TableName() string {
	return "users"
}

var UserDao UserDaoIF

//go:generate mockgen -source user.go  --destination /mocks/user_mock.go --package dao
type UserDaoIF interface {
	AutoMigrate()
	Create(user *User) error
	Update(field string, value interface{}, where string, args ...interface{}) error
	Delete(user *User) error
	Query(where string, args ...interface{}) ([]User, error)
	QueryOrder(user *User, limit int, offset int, orderType string, where string, args ...interface{}) ([]User, error)
	QueryOne(where string, args ...interface{}) (User, error)
	Updates(project *User, where string, args ...interface{}) error
}

type userDao struct {
	client db.Client
}

func NewUserDao(db db.Client) UserDaoIF {
	UserDao = userDao{client: db}
	// 自动建表
	UserDao.AutoMigrate()
	return UserDao
}

func (u userDao) AutoMigrate() {
	u.client.DB().AutoMigrate(&User{})
}

func (u userDao) Create(user *User) error {
	d := u.client.DB().Create(user)
	return d.Error
}

func (u userDao) Update(field string, value interface{}, where string, args ...interface{}) error {
	//u.client.DB().
	result := u.client.DB().Model(&User{}).Where(where, args).Update(field, value)
	return result.Error
}

func (u userDao) Delete(user *User) error {
	panic("implement me")
}

func (u userDao) Query(where string, args ...interface{}) ([]User, error) {
	var users []User
	result := u.client.DB().Find(&users, where, args)
	return users, result.Error
}

func (u userDao) QueryOne(where string, args ...interface{}) (User, error) {
	var user User
	result := u.client.DB().First(&user, where, args)
	return user, result.Error
}
func (u userDao) QueryOrder(project *User, limit int, offset int, orderType string, where string, args ...interface{}) ([]User, error) {
	var projects []User
	db := u.client.DB()
	switch orderType {
	case "time":
		db = db.Order("created_at desc")
	case "score":
		db = db.Order("score desc")
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

func (u userDao) Updates(project *User, where string, args ...interface{}) error {
	result := u.client.DB().Model(&User{}).Where(where, args).Updates(project)
	return result.Error
}
