package router

import (
	"Evo/ctrl"
	"Evo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.POST("/manager/login", ctrl.AdminLogin)
	r.POST("/manager/team", middleware.AuthMW(), ctrl.PostTeam)
	// r.GET("/upload", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK,"upload.html",nil)
	// })
	r.POST("/manager/image",ctrl.PostImage)
	r.POST("/team/login", ctrl.TeamLogin)
	return r
}
