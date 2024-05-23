package outlay

import (
	outlay_base "financialrecord-backend/bases/outlay"
	"financialrecord-backend/dtos/outlay"
	"financialrecord-backend/dtos/pagedlist"
	outlay_model "financialrecord-backend/models/outlay"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TABLE = "outlays"
const NAME = "支出"

func FilterList(tx *gorm.DB, dto outlay.ListDto) *gorm.DB {
	if dto.CatId != nil {
		tx = tx.Where("cat_id = ?", dto.CatId)
	}
	if dto.LowMoney != nil {
		tx.Where("money >= ?", dto.LowMoney)
	}
	if dto.TopMoney != nil {
		tx.Where("money <= ?", dto.TopMoney)
	}
	if dto.STime != nil {
		tx.Where("time >= ?", dto.STime)
	}
	if dto.ETime != nil {
		tx.Where("time < ?", dto.ETime)
	}
	if dto.UserId != nil {
		tx = tx.Where("user_id = ?", dto.UserId)
	}
	return tx
}
func List(db *gorm.DB, dto outlay.ListDto) ([]outlay.Dto, int8, error) {
	var res []outlay.Dto
	tx := db.Model(&outlay_model.Outlay{}).Preload("Cat")
	tx = FilterList(tx, dto).Order("time desc").Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlay.Dto, 0), 1, tx.Error
	}
	return res, 0, nil
}
func Count(db *gorm.DB, dto outlay.ListDto) (int, int8, error) {
	var res int64
	tx := db.Model(&outlay_model.Outlay{})
	tx = FilterList(tx, dto).Count(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表数量失败。%s\n", NAME, tx.Error)
		return 0, 1, tx.Error
	}
	return int(res), 0, nil
}
func PagedList(db *gorm.DB, dto outlay.PagedListDto) (pagedlist.Dto[outlay.Dto], int8, error) {
	res := pagedlist.Empty[outlay.Dto]()
	total, code, err := Count(db, dto.ListDto)
	if code != 0 {
		return res, 1, err
	}
	if total == 0 {
		return res, 0, nil
	}
	tx := db.Model(&outlay_model.Outlay{}).Preload("Cat")
	tx = FilterList(tx, dto.ListDto).Order("time desc").Offset(dto.Index * dto.Size).Limit(dto.Size).Find(&res.Items)
	if tx.Error != nil {
		log.Printf("获取%s分页列表失败。%s\n", NAME, tx.Error)
		return res, 2, tx.Error
	}
	res.Total = total
	return res, 0, nil
}
func Get(db *gorm.DB, id uuid.UUID) (*outlay.Dto, int8, error) {
	var res outlay.Dto
	filter := outlay_model.Outlay{Base: outlay_base.Base{Id: id}}
	tx := db.Model(&filter).Preload("Cat").First(&res, &filter)
	if tx.Error != nil {
		log.Printf("获取%s失败。%s\n", NAME, tx.Error)
		return nil, 1, tx.Error
	}
	return &res, 0, nil
}
func Create(db *gorm.DB, dto outlay.CreateDto) (*uuid.UUID, int8, error) {
	model := outlay_model.Outlay{
		Base: outlay_base.Base{
			Id:    uuid.New(),
			Props: dto,
		},
	}
	tx := db.Create(&model)
	if tx.Error != nil {
		log.Printf("创建%s失败。%s\n", NAME, tx.Error)
		return nil, 1, tx.Error
	}
	return &model.Id, 0, nil
}
func Update(db *gorm.DB, dto outlay.UpdateDto) (bool, int8, error) {
	model := outlay_model.Outlay{
		Base: dto,
	}
	tx := db.Omit("Id").Updates(&model)
	if tx.Error != nil {
		log.Printf("更新%s失败。%s\n", NAME, tx.Error)
		return false, 1, tx.Error
	}
	return true, 0, nil
}
func Delete(db *gorm.DB, id uuid.UUID) (bool, int8, error) {
	tx := db.Delete(&outlay_model.Outlay{Base: outlay_base.Base{Id: id}})
	if tx.Error != nil {
		log.Printf("删除%s失败。%s\n", NAME, tx.Error)
		return false, 1, tx.Error
	}
	return true, 0, nil
}
