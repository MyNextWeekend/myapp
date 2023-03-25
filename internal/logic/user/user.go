package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"myapp/internal/consts"
	"myapp/internal/dao"
	"myapp/internal/model"
	"myapp/internal/model/do"
	"myapp/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
	count, err := dao.UserInfo.Ctx(ctx).Count(do.UserInfo{
		UserName: in.UserName,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count > 0 {
		g.Log().Warningf(ctx, "账号%v已存在", in.UserName)
		return errors.New("账号已存在！！")
	}
	_, err = dao.UserInfo.Ctx(ctx).Insert(do.UserInfo{
		UserName: in.UserName,
		NickName: in.NickName,
		Password: in.Password,
		Role:     "admin",
	})
	return err
}

func (s sUser) Update(ctx context.Context, in model.UserUpdateInput) error {
	count, err := dao.UserInfo.Ctx(ctx).Count(do.UserInfo{
		UserName: in.UserName,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count > 0 {
		g.Log().Infof(ctx, "账号【%v】已存在", in.UserName)
		return errors.New("账号被占用，请更换")
	}
	_, err = dao.UserInfo.Ctx(ctx).Where(do.UserInfo{
		Id: ctx.Value(consts.CtxAdminId)}).Update(do.UserInfo{
		UserName: in.UserName,
		NickName: in.NickName,
		Password: in.Password,
	})
	if err != nil {
		g.Log().Warningf(ctx, "修改用户信息错误：%v", err)
		return err
	}
	g.Log().Infof(ctx, "修改用户【%v】信息为【%v】成功", ctx.Value(consts.CtxUserName), in.UserName)
	return nil
}

func (s sUser) Delete(ctx context.Context) error {
	count, err := dao.UserInfo.Ctx(ctx).Count(do.UserInfo{
		Id:       ctx.Value(consts.CtxAdminId),
		UserName: ctx.Value(consts.CtxUserName),
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count == 0 {
		g.Log().Warningf(ctx, "账号【%v】不存在", ctx.Value(consts.CtxUserName))
		return errors.New("账号不存在")
	}
	_, err = dao.UserInfo.Ctx(ctx).Delete(do.UserInfo{
		Id:       ctx.Value(consts.CtxAdminId),
		UserName: ctx.Value(consts.CtxUserName),
	})
	if err != nil {
		g.Log().Warningf(ctx, "账号【%v】删除失败：%v", ctx.Value(consts.CtxUserName), err)
		return err
	}
	g.Log().Infof(ctx, "【%v】账号删除成功", ctx.Value(consts.CtxUserName))
	return nil
}
