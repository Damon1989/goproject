# 阶段1：编译阶段
FROM golang:1.21-alpine AS builder


# 设置工作目录
WORKDIR /app

# 启用 Go Modules，配置国内代理，禁用 CGO 生成静态二进制
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 复制依赖文件并下载
COPY go.mod go.sum* ./
RUN go mod download

# 复制所有代码
COPY . .

# 编译应用
RUN go build -ldflags "-s -w" -o app main.go

# 阶段2：运行阶段
FROM alpine

# 设置时区
RUN apk --no-cache add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/app .

# 暴露端口
EXPOSE 8090

# 启动应用
CMD ["./app"]