package http

import (
	"github.com/gin-gonic/gin"

	"github.com/skyris/auth-server/internal/env"
)

func (d *Controller) initRouter() {
	if env.APP_ENV != "prod" || env.LOG_LEVEL == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	d.routerAuth(router.Group("/auth"))
	d.router = router
}

func (d *Controller) routerAuth(router *gin.RouterGroup) {
	router.GET("/ping", d.Ping)
	router.POST("/register", d.Register)
	router.POST("/login", d.Login)
	router.GET("/logout", d.Logout)
	router.GET("/welcome", d.Welcome)
	router.GET("/refresh", d.Refresh)
}
