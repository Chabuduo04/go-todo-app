package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Chabuduo04/go-todo-app/internal/controllers"
    "github.com/Chabuduo04/go-todo-app/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/todos", controllers.GetTodos)
        auth.POST("/todos", controllers.CreateTodo)
        auth.PUT("/todos/:id", controllers.UpdateTodo)
        auth.DELETE("/todos/:id", controllers.DeleteTodo)
    }
}
