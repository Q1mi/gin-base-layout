package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/q1mi/gin-base-layout/internal/handler/calc"
	"github.com/q1mi/gin-base-layout/pkg/logging"
	"github.com/spf13/viper"
)

func SetupRoutes(cfg *viper.Viper) *gin.Engine {
	r := gin.New()
	r.Use(logging.GinLogger(), logging.GinRecovery(true)) // 日志中间件，记录请求日志
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Authorization")
	corsCfg.AllowAllOrigins = true
	r.Use(cors.New(corsCfg)) // CORS 跨域中间件，简单粗暴，直接放行所有跨域请求
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/add", calc.AddHandler())
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
