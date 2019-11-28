package router

import (
	"blog/app/api"
	"github.com/gin-gonic/gin"
)

func BlogRouter(r *gin.Engine){
	r.POST("/registeruser", api.RegisterUser)
	r.POST("/login")
	r.POST("/verify")
	r.GET("/refresh")
	r.GET("/user/:name")

}

