package controllers

import (
    "net/http"
    "github.com/Chabuduo04/go-todo-app/pkg/database"
    "github.com/Chabuduo04/go-todo-app/internal/models"
	"github.com/Chabuduo04/go-todo-app/internal/utils"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // 密码加密
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encrypt password"})
        return
    }
    user.Password = string(hashedPwd)

    // 保存用户
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exists or db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func Login(c *gin.Context) {
    var input models.User
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // 查询用户
    if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    // 验证密码
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

	token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
