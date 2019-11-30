package router

import (
	"blog/app/api"
	"github.com/gin-gonic/gin"
)

func BlogRouter(r *gin.Engine) {
	r.POST("/registeruser", api.RegisterUser)
	r.POST("/login", api.Login)
	r.POST("/verify")

	user := r.Group("/user")
	{
		user.GET("/:name")


	}

}
