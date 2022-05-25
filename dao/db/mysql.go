package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NewMysqlDB 初始化mysql驱动
func NewMysqlDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 fmt.Sprintf("%v:%v", "127.0.0.1", "3306"), //IP:PORT
		Net:                  "tcp",
		DBName:               "project",
		Loc:                  time.Local,
		ParseTime:            true,
		AllowNativePasswords: true,
		Params:               map[string]string{"charset": "utf8mb4", "interpolateParams": "true"},
	}
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
