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
// go test (-short) (-bench=".*"):执行测试用例):执行以_test.go结尾的测试用例

func main() {
	server := gin.Default()
	// 静态路由
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 参数路由
	server.GET("/manga/:name", func(ctx *gin.Context) {
		name := ctx.Param("name") // 取得路由参数
		ctx.String(http.StatusOK, "路由参数: "+name)
	})
	// 通配符路由
	server.GET("/page/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "访问的页面: "+ctx.Param(".html"))
	})

	server.GET("/user", func(ctx *gin.Context) {
		id := ctx.Query("id") // 取得请求参数
		ctx.String(http.StatusOK, "用户ID: "+id)
	})

	// *不能单独使用,同时要放在路径尾部
	//server.GET("/wiki/*", func(ctx *gin.Context) {})

	server.Run(":2333")
}
