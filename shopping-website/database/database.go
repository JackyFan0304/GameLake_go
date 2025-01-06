package database

import (
	"log"
	"shopping-website/models" // 確保導入 models 包以進行自動遷移

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局變量，用於存儲數據庫連接

// InitializeDB 初始化數據庫連接
func InitializeDB(dataSourceName string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// 自動遷移，創建表格
	err = DB.AutoMigrate(&models.User{}) // 確保您導入了 models 包以進行自動遷移
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
