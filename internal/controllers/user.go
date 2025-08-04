package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Register(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "register ok"})
}

func Login(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "login ok"})
}
