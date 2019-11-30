package entity

// 自己加引号
// json的规范一般是用snake下划线命名
type User struct {
	Id         string `json:"Userid"`
	Name       string `json:"Username"`
	Gender     string `json:"UserGender"`
	Pwd        string `json:"Password"`
	Permission string `json:"Permission"`
	Phone      string `json:"Phone"`
}

// 看不到黄颜色吗
type RegisterInfo struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

type LoginInfo struct {
	Token string `json:"token"`
}
