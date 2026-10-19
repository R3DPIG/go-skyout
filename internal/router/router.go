package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// New 创建 Gin 引擎并注册路由
func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "pong"})
	})

	// 管理端
	admin := r.Group("/admin")
	_ = admin // TODO: 注册员工登录、菜品、套餐、订单等接口

	// 用户端
	user := r.Group("/user")
	_ = user // TODO: 注册小程序登录、地址簿、购物车、下单等接口

	return r
}
