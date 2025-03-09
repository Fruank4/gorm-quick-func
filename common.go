package main

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WhereFunction func(db *gorm.DB) *gorm.DB

type data struct{}

var ErrDBForbiddenWithoutConstraints = errors.New("db forbidden using constraints not allowed")

func CombineWheresDB(ctx context.Context, functions ...WhereFunction) *gorm.DB {
	db := DB(ctx)
	for _, fc := range functions {
		db = fc(db)
	}
	return db
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

/**
- 带有wheres的update，delete，一定要判断wheres条件是否为空，避免进行全表扫描
- dbRes中会返回rowAffected字段，表示更新的条数，这个一般都要注意一下，是不是更新成功了
*/

// DeleteData 物理删除数据
func DeleteData(ctx context.Context, value interface{}, wheres ...WhereFunction) (int64, error) {
	if len(wheres) == 0 {
		return 0, ErrDBForbiddenWithoutConstraints
	}
	dbRes := CombineWheresDB(ctx, wheres...).Model(value).Delete(value)
	if dbRes.Error != nil {
		return 0, dbRes.Error
	}
	return dbRes.RowsAffected, nil
}

/**
- Update			❌ 只能更新一个字段	❌ 不忽略零值	只更新单个字段（不通用）
- Updates (struct)	✅ 更新多个字段		✅ 忽略零值		适用于更新所有字段
- Updates (map)		✅ 更新多个字段		❌ 不忽略零值	适用于部分字段更新
*/

// BatchUpdateData 更新模型所有字段（零值不更新）
func BatchUpdateData(ctx context.Context, value interface{}, wheres ...WhereFunction) (int64, error) {
	if len(wheres) == 0 {
		return 0, ErrDBForbiddenWithoutConstraints
	}
	db := CombineWheresDB(ctx, wheres...).Model(value)
	dbRes := db.Updates(value)
	if dbRes.Error != nil {
		return 0, dbRes.Error
	}
	return dbRes.RowsAffected, nil
}

/**
- 更新部分字段，控制时可以使用 Select("字段名") 或 Omit("字段名") 。
*/

// UpdateData 更新模型部分字段（零值不更新）
func UpdateData(ctx context.Context, value interface{}, include []string, exclude []string, wheres ...WhereFunction) (int64, error) {
	if len(wheres) == 0 {
		return 0, ErrDBForbiddenWithoutConstraints
	}
	db := CombineWheresDB(ctx, wheres...).Model(value)
	if len(include) > 0 {
		db = db.Select(include)
	}
	if len(exclude) > 0 {
		db = db.Omit(exclude...)
	}
	dbRes := db.Updates(value)
	if dbRes.Error != nil {
		return 0, dbRes.Error
	}
	return dbRes.RowsAffected, nil
}

/**
- 如果想更新零值的所有字段，可以使用Save，部分字段的话，我们可以使用UpdatesWithFields
*/

// UpdateDataWithFields 更新模型的部分字段（零值更新）
func UpdateDataWithFields(ctx context.Context, model interface{}, updateFields map[string]interface{}, wheres ...WhereFunction) (int64, error) {
	if len(wheres) == 0 {
		return 0, ErrDBForbiddenWithoutConstraints
	}
	dbRes := CombineWheresDB(ctx, wheres...).Model(model).Updates(updateFields)
	if dbRes.Error != nil {
		return 0, dbRes.Error
	}
	return dbRes.RowsAffected, nil
}

func FindOne(ctx context.Context, value interface{}, wheres ...WhereFunction) (bool, error) {
	if len(wheres) == 0 {
		return false, ErrDBForbiddenWithoutConstraints
	}
	dbRes := CombineWheresDB(ctx, wheres...).Model(value).First(value)
	if dbRes.Error != nil {
		if errors.Is(dbRes.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, dbRes.Error
	}
	return true, nil
}

func FindData(ctx context.Context, value interface{}, wheres ...WhereFunction) ([]*data, error) {
	var d []*data
	if len(wheres) == 0 {
		return nil, ErrDBForbiddenWithoutConstraints
	}
	err := CombineWheresDB(ctx, wheres...).Model(value).Find(&d).Error
	if err != nil {
		return nil, err
	}
	return d, nil
}

func FindDataCount(ctx context.Context, value interface{}, wheres ...WhereFunction) (int64, error) {
	var count int64
	if len(wheres) == 0 {
		return 0, ErrDBForbiddenWithoutConstraints
	}
	err := CombineWheresDB(ctx, wheres...).Model(value).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
