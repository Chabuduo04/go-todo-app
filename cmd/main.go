package main

import (
	_ "github.com/Chabuduo04/go-todo-app/docs"
	"github.com/Chabuduo04/go-todo-app/internal/routes"
	"github.com/Chabuduo04/go-todo-app/pkg/database"
	"github.com/Chabuduo04/go-todo-app/internal/config"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version 1.0
// @description 一个简易的任务管理系统，支持用户注册、登录和 Todo 管理
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化配置
	config.InitConfig()
	database.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080") // 启动服务
}
