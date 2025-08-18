package controllers

import (
	"net/http"

	"github.com/Chabuduo04/go-todo-app/internal/models"
	"github.com/Chabuduo04/go-todo-app/pkg/database"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Summary 创建 Todo
// @Description 创建新的 Todo 任务
// @Tags todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo 信息"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /todos [post]
// @Security ApiKeyAuth
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

// GetTodo godoc
// @Summary 获取 Todo
// @Description 获取 Todo 任务
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos [get]
// @Security ApiKeyAuth
func GetTodos(c *gin.Context) {
    var todos []models.Todo
    userID := c.MustGet("userID").(uint)
    if err := database.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// UpdateTodo godoc
// @Summary 更新 Todo
// @Description 更新指定 ID 的 Todo 任务
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Todo 信息"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [put]
// @Security ApiKeyAuth
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

// DeleteTodo godoc
// @Summary 删除 Todo
// @Description 删除指定 ID 的 Todo 任务
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [delete]
// @Security ApiKeyAuth
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(uint)
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
