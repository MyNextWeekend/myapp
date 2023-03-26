package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "myapp/api/v1"
	"myapp/internal/consts"
	"myapp/internal/dao"
	"myapp/internal/model/do"
	"myapp/internal/model/entity"
	"myapp/utility"
	"strconv"
)

var gfToken *gtoken.GfToken

/*
NewGToken 创建gToken对象
*/
func NewGToken(ctx context.Context) *gtoken.GfToken {

	gfToken = &gtoken.GfToken{
		CacheMode:        utility.CfgGet(ctx, "gToken.CacheMode").Int8(),
		CacheKey:         utility.CfgGet(ctx, "gToken.CacheKey").String(),
		Timeout:          utility.CfgGet(ctx, "gToken.Timeout").Int(),
		MaxRefresh:       utility.CfgGet(ctx, "gToken.MaxRefresh").Int(),
		TokenDelimiter:   utility.CfgGet(ctx, "gToken.TokenDelimiter").String(),
		EncryptKey:       utility.CfgGet(ctx, "gToken.EncryptKey").Bytes(),
		AuthFailMsg:      utility.CfgGet(ctx, "gToken.AuthFailMsg").String(),
		MultiLogin:       utility.CfgGet(ctx, "gToken.MultiLogin").Bool(),
		LoginPath:        "/login",
		LogoutPath:       "/logout",
		AuthExcludePaths: g.SliceStr{"/register"},
		LoginBeforeFunc:  loginBeforeFunc, //登录时调用
		LoginAfterFunc:   loginAfterFunc,  //登录后调用
		AuthAfterFunc:    authAfterFunc,   //拦截路由调用
	}
	return gfToken
}

/*
登录校验生成token
*/
func loginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	userName := r.Get("userName").String()
	passwd := r.Get("password").String()
	if len(userName) == 0 || len(passwd) == 0 {
		r.Response.WriteJson(gtoken.Fail("账号或密码不能为空."))
		r.ExitAll()
	}
	var user *entity.UserInfo
	err := dao.UserInfo.Ctx(r.GetCtx()).Where(do.UserInfo{UserName: userName, Password: passwd}).Scan(&user)
	if err != nil {
		g.Log().Info(r.GetCtx(), "查询数据库失败", err)
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	if user == nil {
		g.Log().Info(r.GetCtx(), "查询结果为空", err)
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	if passwd != user.Password {
		g.Log().Info(r.GetCtx(), "账号密码错误", err)
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	g.Log().Info(r.GetCtx(), user.UserName, "登录成功")

	//第一个字段是用户唯一标识，第二个字段是扩展参数
	return strconv.Itoa(int(user.Id)), user
}

/*
根据登录成功的token，组装用户返回信息
*/
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
		r.ExitAll()
	}
	//获取用户的id
	userKey := respData.GetInt("userKey")
	//根据用户的id获取用户的其他信息
	var user *entity.UserInfo
	err := dao.UserInfo.Ctx(r.GetCtx()).Where(do.UserInfo{Id: userKey}).Scan(&user)
	if err != nil {
		g.Log().Info(r.GetCtx(), "查询数据库失败", err)
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	respData.Data = &v1.UserLoginRes{
		UserName:   user.UserName,
		Role:       user.Role,
		CreateTime: user.CreatedAt,
		UpdateTime: user.UpdatedAt,
		Token:      respData.GetString("token"),
		Type:       "Bearer",
	}
	r.Response.WriteJson(respData)
}

/*
拦截路由操作,根据token解析情况，设置上下文数据
*/
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//判断token解析的状态
	if !respData.Success() {
		r.Response.WriteJson(respData)
		r.ExitAll()
	}
	//将token信息转换成对象，便于后续取数据
	var data *entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &data)
	if err != nil {
		g.Log().Error(r.GetCtx(), "结构体转换失败：%v", err)
		panic(err)
	}
	//将对象信息写入上下文
	r.SetCtxVar(consts.CtxAdminId, data.Id)
	r.SetCtxVar(consts.CtxUserName, data.UserName)
	r.SetCtxVar(consts.CtxRole, data.Role)
	r.Middleware.Next()
}
