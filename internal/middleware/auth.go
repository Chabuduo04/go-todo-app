package middleware

import (
    "net/http"
    "strings"
    "github.com/Chabuduo04/go-todo-app/internal/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            c.Abort()
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ParseJWT(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
