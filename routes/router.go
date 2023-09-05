package routes

import (
	api "eletronicMall/api/v1"
	"eletronicMall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 跨域 中间件
	router.Use(middleware.Cors())
	router.StaticFS("/static", http.Dir("./static")) // 静态文件加载
	v1 := router.Group("/api/v1")                    // 路由分组
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.POST("user/register", api.Register) // 注册
		v1.POST("user/login", api.Login)       // 登录
		// 需要登录保护的
		authentication := v1.Group("/")
		authentication.Use(middleware.JWTAuth())
		{
			// 用户操作
			authentication.PUT("user", api.UserUpdate)
			authentication.POST("avatar", api.UserUploadAvatar)

			// 邮箱操作
			authentication.POST("user/sendemail", api.SendEmail)
		}
	}

	return router
}
