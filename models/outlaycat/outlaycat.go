package outlaycat

import (
	"financialrecord-backend/bases/outlay"
	"financialrecord-backend/bases/outlaycat"

	"gorm.io/gorm"
)

// 支出类型
type OutlayCat struct {
	outlaycat.Base
	Parent      *outlaycat.Base  `gorm:"foreignKey:Id;references:ParentId" json:"parent"`   // 父级
	Children    []outlaycat.Base `gorm:"foreignKey:ParentId;references:Id" json:"children"` // 子级列表
	Outlays     []outlay.Base    `gorm:"foreignKey:CatId;references:Id" json:"outlays"`     // 支出列表
	HasChildren bool             `gorm:"-:all" json:"hasChildren"`
}

func (cat *OutlayCat) AfterFind(tx *gorm.DB) (err error) {
	res := int64(0)
	model := OutlayCat{Base: outlaycat.Base{
		Props: outlaycat.Props{
			ParentId: cat.Id,
		},
	}}
	tx.Model(&model).Where(&model).Count(&res)
	cat.HasChildren = res > 0
	return
}

func (m *OutlayCat) TableName() string {
	return "outlay_cats"
}
