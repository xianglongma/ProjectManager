package main

import (
	"fmt"
	"github.com/xianglongma/ProjectManager/dao/db"
	"github.com/xianglongma/ProjectManager/server"
)

// 服务端程序入口
func main() {
	fmt.Println("项目管理与评价网站开始开发")
	// 初始化操作
	Init()
	// run http server
	RunServer()
}

func Init() error {
	// 初始化数据库
	if err := db.InitDBClient(); err != nil {
		return err
	}
	return nil
}

func RunServer() {
	server.NewGinServer("0.0.0.0:80").SetRouter().RunServer()
}
