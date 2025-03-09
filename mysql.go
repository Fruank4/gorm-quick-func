package main

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	gormDB *gorm.DB
	once   sync.Once
)

func Init() {
	once.Do(func() {
		dsn := "root:123456@tcp(127.0.0.1:3306)/test_db_name?charset=utf8&parseTime=True&loc=Local"
		// 实际上，Mysql.Config 还可以使用更多的配置，比如字段的默认长度，datetime的精度，根据当前Mysql版本自动配置
		// 以及使用其他类型的数据库：https://gorm.io/zh_CN/docs/connecting_to_the_database.html
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		gormDB = db
		gormDB = DB.Debug() // 打开db的debug模式，用于显示logç
	})

}

func WriteDB(ctx context.Context) *gorm.DB {
	return gormDB.Clauses().WithContext(ctx)
}

func ReadDB(ctx context.Context) *gorm.DB {
	return gormDB.Clauses().WithContext(ctx)
}

func DB(ctx context.Context) *gorm.DB {
	return gormDB.WithContext(ctx)
}
