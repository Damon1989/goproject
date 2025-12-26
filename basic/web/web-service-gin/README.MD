go linux 无法下载第三方包时的解决办法：
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct



docker build -t go-gin-app:v1 .
docker run -d -p 8090:8090 --name go-gin-app-container go-gin-app:v1
