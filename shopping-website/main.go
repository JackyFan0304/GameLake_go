package main

import (
	"log"
	"os"
	"shopping-website/controllers"
	"shopping-website/database"
	"shopping-website/models"
	"shopping-website/routes"
	"shopping-website/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化數據庫連接
	// 加載 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dataSourceName := os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("資料庫連線" + dataSourceName)
	database.InitializeDB(dataSourceName)

	// 初始化資料庫連接
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// 自動遷移，根據模型創建資料表
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully!")

	// 創建 GIN 路由
	r := gin.Default()

	// 初始化服務
	userModel := models.NewUserModel(database.DB)                        // 使用 UserModel                                     // 確保您有 UserModel 的實現
	authService := services.NewAuthService(userModel, "your_secret_key") // 替換為您的 JWT 密鑰

	// 初始化控制器
	authController := controllers.NewAuthController(authService)

	// 設置路由
	authRoutes := r.Group("/auth")                     // 創建 /auth 路由組
	routes.SetupAuthRoutes(authRoutes, authController) // 傳遞 authRoutes 和 authController

	// 啟動伺服器
	r.Run(":8080") // Listen and serve on port 8080
}
