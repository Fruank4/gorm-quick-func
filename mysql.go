package main

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func WriteDB(ctx context.Context) *gorm.DB {
	return db.Clauses().WithContext(ctx)
}

func ReadDB(ctx context.Context) *gorm.DB {
	return db.Clauses().WithContext(ctx)
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
