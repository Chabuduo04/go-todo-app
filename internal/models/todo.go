package models

import (
    "gorm.io/gorm"
)

type Todo struct {
    gorm.Model  // 包含ID, CreatedAt, UpdatedAt, DeletedAt字段
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
    UserID    uint   `json:"user_id"`
}
