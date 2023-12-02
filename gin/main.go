package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// go env:查看环境变量
// go list -m all:查看已安装的包
// go list -m -versions github.com/gin-gonic/gin:查看可用版本
// go get github.com/gin-gonic/gin@v1.8.0:安装指定版本
// go mod tidy:同步依赖
// go get (-u):更新依赖
// go test:执行以_test.go结尾的测试用例

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
