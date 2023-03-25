package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 登录之后的响应体
type UserLoginRes struct {
	UserName   string
	Role       string
	CreateTime *gtime.Time
	UpdateTime *gtime.Time
	Token      string
	Type       string
}

type UserRegisterReq struct {
	g.Meta    `path:"/user/register" method:"post" summary:"注册账号" tags:"用户"`
	UserName  string `json:"userName" v:"required|length:4,30#请输入账号|账号长度为:min到:max位" dc:"账号"`
	NickName  string `json:"nickName" dc:"昵称"`
	Password  string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够" dc:"密码"`
	Password2 string `p:"password2" v:"required|length:6,30|same:Password#请确认密码|密码长度不够|两次密码不一致" dc:"密码"`
}

type UserRegisterRes struct {
	Ok bool
}

type UserUpdateReq struct {
	g.Meta    `path:"/user/update" method:"post" summary:"修改信息" tags:"用户"`
	UserName  string `json:"userName" v:"required|length:4,30#请输入账号|账号长度为:min到:max位" dc:"账号"`
	NickName  string `json:"nickName" dc:"昵称"`
	Password  string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够" dc:"密码"`
	Password2 string `p:"password2" v:"required|length:6,30|same:Password#请确认密码|密码长度不够|两次密码不一致" dc:"密码"`
}

type UserUpdateRes struct {
	Ok bool
}
type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" summary:"删除用户" tags:"用户"`
}

type UserDeleteRes struct {
	Ok bool
}
