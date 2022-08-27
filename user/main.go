package main

import (
	"fmt"
	"log"
	"user/conf"
	"user/model"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Config.Mysql.User, conf.Config.Mysql.Password, conf.Config.Mysql.Host, conf.Config.Mysql.Port, conf.Config.Mysql.DbName)
	if model.Database(dsn) != nil {
		log.Fatal("迁移数据库失败")
	} else {
		fmt.Println("迁移成功")
	}
}
