package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库单例
var DB Client

type Client interface {
	DB() *gorm.DB
}

type MysqlClient struct {
	db *gorm.DB
}

func (c *MysqlClient) DB() *gorm.DB {
	return c.db
}

func NewMysqlClient() (Client, error) {
	db, err := NewMysqlDB()
	if err != nil {
		return nil, nil
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	return &MysqlClient{db: gormDB}, err
}

func InitDBClient() error {
	db, err := NewMysqlClient()
	if err != nil {
		return nil
	}
	DB = db
	return nil
}
