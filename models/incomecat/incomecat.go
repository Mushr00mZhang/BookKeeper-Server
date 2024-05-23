package incomecat

import "github.com/google/uuid"

type Base struct {
	ParentId *uuid.UUID `gorm:""`                  // 父级Id
	Name     string     `gorm:"not null;size:64;"` // 名称
	Unit     string     `gorm:"not null;size:16;"` // 单位
	Sort     uint8      `gorm:"not null;"`         // 排序
	Stable   bool       `gorm:"not null;"`         // 是否固定
	Remark   *string    `gorm:"size:256;"`         // 备注
}

// 收入类型
type IncomeCat struct {
	Base
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
}
