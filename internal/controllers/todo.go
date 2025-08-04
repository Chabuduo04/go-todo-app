package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetTodos(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "get todos"})
}

func CreateTodo(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "create todo"})
}

func UpdateTodo(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "update todo"})
}

func DeleteTodo(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "delete todo"})
}
