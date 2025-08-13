package controllers

import (
    "net/http"
    "github.com/Chabuduo04/go-todo-app/pkg/database"
    "github.com/Chabuduo04/go-todo-app/internal/models"

    "github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    userID := c.MustGet("userID").(uint)
    todo.UserID = userID

    if err := database.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "todo created successfully", "todo": todo})
}

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    userID := c.MustGet("userID").(uint)
    if err := database.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    userID := c.MustGet("userID").(uint)
    var todo models.Todo

    if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
        return
    }

    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    if err := database.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "todo updated successfully", "todo": todo})
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    userID := c.MustGet("userID").(uint)
    if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
