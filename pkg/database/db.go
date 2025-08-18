package database

import (
    "fmt"
    "log"

    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/Chabuduo04/go-todo-app/internal/models"
    "github.com/Chabuduo04/go-todo-app/internal/config"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.AppConfig.DBUser,
        config.AppConfig.DBPass,
        config.AppConfig.DBHost,
        config.AppConfig.DBPort,
        config.AppConfig.DBName,
    )
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
