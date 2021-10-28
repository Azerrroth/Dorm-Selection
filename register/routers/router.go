package routers

import (
	"register/pkg/setting"
	v1 "register/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	// r.Use(sessions.Sessions("session", util.Store))

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	// apiv1.POST("/login", v1.Login)
	apiv1.POST("/register", v1.Register)

	// apiv1.Use(jwt.JWT())
	// apiv1.Use(session.Session())
	// {
	// 	apiv1.POST("/dorm", v1.GetDormList)
	// }

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
