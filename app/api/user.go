package api

import (
	"blog/app/entity"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context){
	var registerinfo entity.RegisterInfo
	err :=ctx.ShouldBindJSON(&registerinfo)
	if err != nil{
		// 错误原因
		//fmt.Println(registerinfo)
		//_ = ctx.AbortWithError(400, err)
		ctx.JSON(400, gin.H{
			"status": -1,
			"msg": "parameter error",
		})

	}else{
		//检验数据

		//密码加密



	}
	print(registerinfo.Phone)
	print(registerinfo.Password)
	ctx.JSON(200, registerinfo)
}