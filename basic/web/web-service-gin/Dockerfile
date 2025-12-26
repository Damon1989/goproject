# 使用官方的go基础镜像
FROM golang:1.21-alpine as builder

# 设置工作目录
WORKDIR /app

# 开启Go模块（如果你的项目用go mod）
# 国内代理加速依赖下载
# 禁用CGO，编译出静态链接的二进制文件
# 目标系统为Linux
# 目标架构为amd64（根据服务器调整）
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

# 复制go.mod和go.sum（先复制依赖文件，利用Docker缓存，避免每次改代码都重新下载依赖）
COPY go.mod go.sum* ./
RUN go mod download

# 复制所有代码到容器
COPY . .

# 编译应用，输出为app二进制文件
RUN go build -ldflags="-s -w" -o app main.go  # -ldflags="-s -w" 进一步减小二进制体积

# 第二阶段：运行阶段（使用轻量的alpine镜像，仅保留编译后的二进制文件）
FROM alpine

# 可选：添加时区支持（如果应用需要）
RUN apk --no-cache add tzdata
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件到当前镜像
COPY --from=builder /app/app .

# 暴露应用监听的端口
EXPOSE 8090

# 启动应用
CMD ["./app"]