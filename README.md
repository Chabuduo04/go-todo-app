# Go Todo App

[中文 | English](#language-switch)

---

## 中文

一个基于 Go + Gin + MySQL + Redis 的简易任务管理系统，支持用户注册、登录、JWT 权限控制以及 Todo 的增删改查，并已完成容器化部署。本项目适合 Go 新手学习 Web 开发，也适用于小型任务管理系统的快速搭建。

## 功能特性

**用户模块**
-用户注册（密码加密存储）
-用户登录（JWT 认证）

**Todo 模块**
-创建 / 查询 / 更新 / 删除 Todo
-支持用户权限控制（每个用户只能操作自己的 Todo）
-Redis 缓存加速查询

**工程化支持**
-Docker 化 & docker-compose 一键启动
-配置管理（环境变量 & .env 文件）
-Swagger 自动化 API 文档
-单元测试 & 集成测试（go test）

## 技术栈

后端框架: Gin
数据库: MySQL
缓存: Redis
ORM: GORM
认证: JWT (基于 github.com/dgrijalva/jwt-go)
容器化: Docker + docker-compose
API 文档: Swagger (swaggo/gin-swagger)

## 配置

项目依赖 .env 文件来配置环境变量：

```bash
# .env.example
DB_USER=root
DB_PASSWORD=password
DB_HOST=127.0.0.1   #容器化部署改为自己 docker-compose.yml 中的服务名 例如 db
DB_PORT=3306
DB_NAME=todo_db

JWT_SECRET=your_secret_key

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASS=password
```

## 启动方式

### 1. 本地启动

需要本地安装 MySQL & Redis：

```bash
go mod tidy
go run cmd/main.go
```

### 2. Docker Compose 启动（推荐）

```bash
docker-compose up --build
```

启动完成后，服务运行在 http://localhost:8080。

## API 文档

启动服务后访问：

👉 Swagger UI: http://localhost:8080/swagger/index.html

## 未来计划

| 功能 | 状态 | 说明 |
|------|------|------|
| 接口分页 & 搜索 | ⬜ | 待实现 |
| 接口限流（基于 Redis） | ⬜ | 待实现 |
| CI/CD 集成 | ⬜ | 待实现 |
| Kubernetes 部署支持 | ⬜ | 待实现 |
| 监控 & 日志收集 | ⬜ | 待实现 |

## 贡献方式

欢迎提 Issue 或 Pull Request 进行交流与贡献！

## License

MIT License

---

## English

A simple task management system built with Go + Gin + MySQL + Redis.
It supports user registration, login, JWT-based authentication, and CRUD operations for todos.
The project is fully containerized with Docker and suitable for both learning Go web development and building lightweight todo applications.

## Features

**User Module**
-User registration (with password encryption)
-User login (JWT authentication)

**Todo 模块**
-Create / Read / Update / Delete todos
-User permission control (users can only manage their own todos)
-Redis caching to speed up queries

**工程化支持**
-Dockerized & one-click startup with docker-compose
-Config management via environment variables (.env file)
-Swagger for automated API documentation
-Unit tests & integration tests (go test)

## Tech Stack

Framework: Gin
Database: MySQL
Cache: Redis
ORM: GORM
Authentication: JWT (基于 github.com/dgrijalva/jwt-go)
Containerization: Docker + docker-compose
API Docs: Swagger (swaggo/gin-swagger)

## Configuration

The project requires a .env file for environment variables:

```bash
# .env.example
DB_USER=root
DB_PASSWORD=password
DB_HOST=127.0.0.1   # In containerized deployment, replace with the service name in docker-compose.yml, e.g. db
DB_PORT=3306
DB_NAME=todo_db

JWT_SECRET=your_secret_key

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASS=password
```

## Getting Started

### 1. Run locally

Requires MySQL & Redis installed locally:

```bash
go mod tidy
go run cmd/main.go
```

### 2. Run with Docker Compose (recommended)

```bash
docker-compose up --build
```

The app will be available at:
👉 http://localhost:8080

## API Documentation

Once the service is running:

👉 Swagger UI: http://localhost:8080/swagger/index.html

## Roadmap

| Feature | Status | Notes |
|------|------|------|
| Pagination & Search | ⬜ | To be implemented |
| Rate limiting (Redis-based) | ⬜ | To be implemented |
| CI/CD | ⬜ | To be implemented |
| Kubernetes Deployment | ⬜ | To be implemented |
| Monitoring & Logging | ⬜ | To be implemented |

## Contribution

Contributions are welcome! Feel free to open an Issue or submit a Pull Request.

## License

MIT License

---

## <span id="language-switch"></span>Language Switch

- [中文](#中文)
- [English](#english)
