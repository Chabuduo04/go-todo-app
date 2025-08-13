package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
    "github.com/Chabuduo04/go-todo-app/internal/models"
)

var DB *gorm.DB

func InitDB() {
    // 读取 .env
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 从环境变量读取
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // 自动迁移
    err = db.AutoMigrate(&models.User{}, &models.Todo{})
    if err != nil {
        log.Fatal("AutoMigrate failed:", err)
    }

    DB = db
}
