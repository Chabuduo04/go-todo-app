package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
    "go-todo-app/internal/models"
)

var DB *gorm.DB

func InitDB() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("MYSQL_DSN")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // ×Ô¶¯Ç¨ÒÆ
    err = db.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatal("AutoMigrate failed:", err)
    }

    DB = db
}
