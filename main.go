package main

import (
	"blog/app/router"
	"github.com/gin-gonic/gin"
)





func main()  {
	r := gin.Default()
	router.BlogRouter(r)
	r.Run()
	
}
