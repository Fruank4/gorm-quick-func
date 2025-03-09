package main

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

/**
Notice：
	1、gorm在底层可以根据value的model进行判断，自动识别表名，因此大多时候不用显示指定Model
*/

/**
- 创建一条新的记录，适用于无脑插入，不存在刷数需求的场景下
- ID如果已经存在，如果自增会自动生成，否则会报错duplicateKey冲突
- 批量插入的时候可以用crateInBatches提高性能
*/
// createData 插入
func createData(ctx context.Context, value interface{}) error {
	return DB(ctx).Create(value).Error
}

// batchCreateData 批量插入
func batchCreateData(ctx context.Context, values interface{}) error {
	return DB(ctx).CreateInBatches(values, 100).Error
}

/**
- 保存一条记录，适用于无法判断这条数据是否已经存在的场景下
- 根据主键判断，如果重复了，则会更新其原有的值（零值也会更新）
- 可以通过Conflict做增强，将判断重复条件扩展到更多的字段上，见下
*/

// SaveData 保存（更新所有值，零值更新）
func SaveData(ctx context.Context, value interface{}) error {
	return DB(ctx).Save(value).Error
}

/**
- columnNames 作为冲突匹配字段（通常是唯一key索引）,如果这些字段有冲突，则执行更新
- updateFields 这些字段作为需要更新的
- 既然已经加了 ON CONFLICT 子句，就可以直接使用 Create，不需要用 Save了。
- 因为Save 会先 SELECT 检查数据是否存在，然后决定 INSERT 还是 UPDATE
- Create+Conflict冲突时自动 UPDATE，不会报错
*/

// SaveDataWithFields 单个保存（自定义更新）
func SaveDataWithFields(ctx context.Context, value interface{}, columnNames []string, updateFields []string) error {
	columns := make([]clause.Column, 0)
	for _, name := range columnNames {
		columns = append(columns, clause.Column{Name: name})
	}
	return DB(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		DoUpdates: clause.AssignmentColumns(updateFields),
		UpdateAll: len(updateFields) > 0,
	}).Create(value).Error
}

/**
- BatchSaveDataWithFields 批量保存（自定义更新）
- Save并没有原生批量的方式，我们可以通过下面这种方式实现
- 如果在代码中就可以判断，我们也可以用手动分区的方式来加速（createInBatches+updates）
*/

// BatchSaveDataWithUpdateFields 批量保存（自定义更新）
func BatchSaveDataWithUpdateFields(ctx context.Context, value interface{}, columnNames []string, updateFields []string) error {
	columns := make([]clause.Column, 0)
	for _, name := range columnNames {
		columns = append(columns, clause.Column{Name: name})
	}
	return DB(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		DoUpdates: clause.AssignmentColumns(updateFields),
		UpdateAll: len(updateFields) > 0,
	}).CreateInBatches(value, 100).Error
}

/**
- UpdateAll/DoUpdates 会尝试更新字段。
- Omit 排除指定的字段。
- 即便同时设置 UpdateAll/DoUpdates 和 Omit，Omit都会优先排除指定的字段，不会被更新。
- 并且Omit 和 DoUpdates 的优先级与它们的调用顺序没有直接关系
*/

// BatchSaveDataWithExcludeFields 批量保存（自定义排除）
func BatchSaveDataWithExcludeFields(ctx context.Context, value interface{}, columnNames []string, excludeFields []string) error {
	columns := make([]clause.Column, 0)
	for _, name := range columnNames {
		columns = append(columns, clause.Column{Name: name})
	}
	return DB(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Omit(excludeFields...).CreateInBatches(value, 100).Error
}
