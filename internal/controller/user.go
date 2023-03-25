package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "myapp/api/v1"
	"myapp/internal/model"
	"myapp/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	g.Log().Info(ctx, "Register接收到的参数是：", req)
	if len(req.NickName) == 0 {
		g.Log().Info(ctx, "由于昵称为空，将账户赋值给昵称")
		req.NickName = req.UserName
	}
	err = service.User().Register(ctx, model.UserRegisterInput{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
	})
	res = &v1.UserRegisterRes{}
	if err != nil {
		res.Ok = false
		return
	}
	res.Ok = true
	return
}

func (c cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	err = service.User().Update(ctx, model.UserUpdateInput{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
	})
	res = &v1.UserUpdateRes{}
	if err != nil {
		res.Ok = false
		return res, err
	}
	res.Ok = true
	return res, err
}

func (c cUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	err = service.User().Delete(ctx)
	res = &v1.UserDeleteRes{}
	if err != nil {
		res.Ok = false
		return
	}
	res.Ok = true
	return
}
