package outlaycat

import "github.com/google/uuid"

// 支出类型属性
type Props struct {
	ParentId uuid.UUID `gorm:"not null;" json:"parentId"`     // 父级Id
	Name     string    `gorm:"not null;size:64;" json:"name"` // 名称
	Unit     string    `gorm:"not null;size:16;" json:"unit"` // 单位
	Sort     uint8     `gorm:"not null;" json:"sort"`         // 排序
	Stable   bool      `gorm:"not null;" json:"stable"`       // 是否固定
	Remark   *string   `gorm:"size:256;" json:"remark"`       // 备注
}

func (m *Props) TableName() string {
	return "outlay_cats"
}

// 支出类型基类
type Base struct {
	Id uuid.UUID `gorm:"not null;primaryKey;references:Id" json:"id"` // Id
	Props
}

func (m *Base) TableName() string {
	return "outlay_cats"
}
