package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	//严格来讲，这不是优秀的实践，应该走审批流程
	return db.AutoMigrate(&User{})
}
