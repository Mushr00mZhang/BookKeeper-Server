package database

import (
	"errors"
	"log"

	"financialrecord-backend/models/outlay"
	"financialrecord-backend/models/outlaycat"

	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化
func Init(db *gorm.DB) error {
	log.Println("数据库初始化中。")
	if db == nil {
		log.Println("数据库为空。")
		return errors.New("db is nil")
	}
	log.Println("数据库迁移中。")
	err := db.AutoMigrate(&outlay.Outlay{}, &outlaycat.OutlayCat{})
	if err != nil {
		log.Println("数据库迁移失败。")
		return err
	}
	log.Println("数据库迁移成功。")
	log.Println("数据库初始化成功。")
	DB = db
	return nil
}
