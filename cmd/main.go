package main

import (
	"fmt"
    "go-todo-app/pkg/database"
    "go-todo-app/internal/routes"
    "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
    database.InitDB()

    r := gin.Default()
    routes.SetupRoutes(r)

	fmt.Println("Server started at http://localhost:8080")
    r.Run(":8080") // Æô¶¯·þÎñ
}
