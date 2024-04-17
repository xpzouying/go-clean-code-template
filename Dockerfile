# 第一阶段：使用 Go 环境进行构建
FROM golang:1.21 AS builder

# 设置工作目录
WORKDIR /app

# 首先只复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 安装依赖并构建项目
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 然后复制项目文件到 Docker 容器
COPY . .


# 构建项目
RUN make build

# RUN ls -la /app/dist/
# RUN find /app

# 第二阶段：运行编译好的二进制文件
FROM alpine:latest

# 把二进制文件从 builder 阶段复制到当前阶段
COPY --from=builder /app/dist/go-template-project /app/go-template

# 需要设定工作区，有保存的相对路径
WORKDIR /app

ENV PORT 8080
EXPOSE 8080

# 运行程序
CMD ["/app/go-template", "-l", "0.0.0.0:8080"]
