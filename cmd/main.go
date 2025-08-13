package main

import (
    "github.com/Chabuduo04/go-todo-app/pkg/database"
    "github.com/Chabuduo04/go-todo-app/internal/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()

    r := gin.Default()
    routes.SetupRoutes(r)

    r.Run(":8080") // 启动服务
}
