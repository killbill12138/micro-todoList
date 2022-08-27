package model

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//数据库链接单例
var DB *gorm.DB

//Database在中间件中初始化mysql链接
func Database(dsn string) error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
	})

	if err != nil {
		return err
	}

	rawDB, err := db.DB()
	if err != nil {
		return err
	}
	rawDB.SetMaxIdleConns(20)
	rawDB.SetMaxOpenConns(100)
	rawDB.SetConnMaxLifetime(time.Second * 45)
	DB = db
	return migration()
}
