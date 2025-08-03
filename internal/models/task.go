package models

type Task struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
    UserID    uint   `json:"user_id"`
}
