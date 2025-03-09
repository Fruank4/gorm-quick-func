package model

import "time"

const tableNameSystemOpLog = "system_op_log"

type SystemOpLog struct {
	ID            int64 `gorm:"column:id;primary_key;autoIncrement:true;comment:自增主键" json:"id"`
	EntityType    int32 `gorm:"column:entity_type;not null;comment:实体类型" json:"entity_type"`
	EntityID      int64
	OperationType int32
	OperatorID    int64
	Content       string
	Extra         *string
	// Gorm约定，使用CreateAt、UpdateAt 追踪创建/更新时间，自动填充 ‘当前时间’，如果想要存储UNIX秒时间戳，而不是time，只需要将time.Time修改为int即可
	CreateTime time.Time
	UpdateTime time.Time
}

func (*SystemOpLog) TableName() string {
	return tableNameSystemOpLog
}
