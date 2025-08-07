package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Chabuduo04/go-todo-app/internal/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    todo := r.Group("/todos")
    {
        todo.GET("", controllers.GetTodos)
        todo.POST("", controllers.CreateTodo)
        todo.PUT("/:id", controllers.UpdateTodo)
        todo.DELETE("/:id", controllers.DeleteTodo)
    }
}
