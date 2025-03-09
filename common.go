package main

import (
	"context"
	"gorm.io/gorm"
)

type WhereFunction func(db *gorm.DB) *gorm.DB

type data struct {
}

func FindData(db *gorm.DB, ctx context.Context, model interface{}, cond WhereFunction) ([]*data, error) {
	var d []*data
	err := cond(DB(ctx)).Model(model).Find(&d).Error
	if err != nil {
		return nil, err
	}
	return d, nil
}
