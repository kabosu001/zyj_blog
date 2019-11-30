package api

import (
	"blog/app/entity"
	"github.com/gin-gonic/gin"
)

//注册
func RegisterUser(ctx *gin.Context) {
	var registerinfo entity.RegisterInfo
	err := ctx.ShouldBindJSON(&registerinfo)
	if err != nil {
		// 错误原因
		//fmt.Println(registerinfo)
		//_ = ctx.AbortWithError(400, err)
		ctx.JSON(400, gin.H{
			"status": -1,
			"msg":    "parameter error",
		})

	} else {
		//检验数据

		//密码加密

		//存mysql

		//主键hashids作为用户id

		//签发jwt

		//302到登录页面

	}

}

//登录
func Login(ctx *gin.Context) {
	var loginReq entity.LoginReq
	err := ctx.BindJSON(&loginReq)
	if err != nil {
		//解析错误
	} else {

	}

}

//检测登录参数
func LoginCheck(){}

//检测手机号是否注册过
func UserPhoneCheck(){}

//用户更新信息
func UserUpdate(){}

//用户注销
func UserDelete(){}

//重新设置密码
func UserPwdReset(){}

