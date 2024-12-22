# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod tidy

# 复制所有文件到容器中
COPY . .

# 构建 Go 应用
RUN go build -o main .

# 设置容器启动时执行的命令
CMD ["./main"]

# 暴露容器端口（假设应用在 8080 端口运行）
EXPOSE 8080
