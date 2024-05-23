package outlay

import (
	"financialrecord-backend/bases/outlay"
	"financialrecord-backend/bases/outlaycat"
)

// 支出
type Outlay struct {
	outlay.Base
	Cat outlaycat.Base `gorm:"foreignKey:CatId;references:Id" json:"cat"` // 类型
}

func (m *Outlay) TableName() string {
	return "outlays"
}
