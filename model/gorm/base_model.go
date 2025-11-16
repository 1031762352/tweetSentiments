package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 公共基础字段
type BaseModel struct {
	ID        uint64         `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:'主键ID'" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type datetime;not null;default CURRENT_TIMESTAMP;comment:'创建时间'" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type datetime;not null;default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type datetime;index;comment:'软删除时间'" json:"-"` // 软删除字段，JSON序列化忽略
}
