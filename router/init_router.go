package router

import "github.com/gin-gonic/gin"

func blogRouter(r *gin.Engine){
	r.POST("/login")
	r.POST("/verify")
	r.GET("/refresh")

}

