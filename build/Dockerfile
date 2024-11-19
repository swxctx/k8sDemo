# 第一阶段：构建应用程序
FROM golang:1.20 AS builder

# 设置工作目录，如果 /app 不存在会自动创建
WORKDIR /app

# 将源代码复制到工作目录中
COPY .. .

# 下载依赖项
RUN go mod download

# 构建 Go 应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mydemo

# 第二阶段：使用 alpine 运行应用程序
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件到 /app 目录
COPY --from=builder /app/mydemo .
COPY --from=builder /app/config/config.yaml ./config/

# 设置容器启动时运行的命令
CMD ["./mydemo"]
