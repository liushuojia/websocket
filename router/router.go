package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api"
)

// 后台中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// 路由初始化
func InitRouter() *gin.Engine {

	router := gin.New()

	// 全局中间件
	router.Use(gin.Recovery())

	// 404
	router.NoRoute(func(c *gin.Context) {
		//返回404状态码
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "404, page not exists!",
		})
		return
	})

	// 服务器状态
	router.GET("/link/:uuid", api.LinkWS)

	return router
}

// 开启http服务
func HttpStart() {
	// http server
	router := InitRouter()
	router.Run(":80")
	return
}
