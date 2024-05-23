package outlay

import (
	"time"

	"github.com/google/uuid"
)

// 支出属性
type Props struct {
	Name     string    `gorm:"not null;size:64;" json:"name"`           // 名称
	CatId    uuid.UUID `gorm:"not null;foreignKey:CatId;" json:"catId"` // 类型Id
	Money    float32   `gorm:"not null;precision:2;" json:"money"`      // 金额
	Original *float32  `gorm:"precision:2;" json:"original"`            // 原价
	Amount   *float32  `gorm:"precision:2;" json:"amount"`              // 数量
	Unit     *string   `gorm:"" json:"unit"`                            // 单位
	Time     time.Time `gorm:"not null;" json:"time"`                   // 时间
	UserId   uuid.UUID `gorm:"not null;" json:"userId"`                 // 用户Id
}

func (m *Props) TableName() string {
	return "outlays"
}

// 支出基类
type Base struct {
	Id uuid.UUID `gorm:"not null;primaryKey;" json:"id"` // Id
	Props
}

func (m *Base) TableName() string {
	return "outlays"
}
