package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model  // 包含ID, CreatedAt, UpdatedAt, DeletedAt字段
    Username string `gorm:"unique" json:"username"`
    Password string `json:"password"`
}
