# 1. 使用官方 Go 镜像
FROM golang:1.24-alpine AS builder

# 2. 设置工作目录
WORKDIR /app

# 3. 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 4. 复制代码
COPY . .

# 5. 编译可执行文件
RUN go build -o todo-app ./cmd/main.go

# 6. 使用更小的镜像运行
FROM alpine:latest

WORKDIR /root/

# 复制可执行文件
COPY --from=builder /app/todo-app .
# 复制 .env
COPY .env .

# 暴露端口
EXPOSE 8080

# 启动服务
CMD ["./todo-app"]
