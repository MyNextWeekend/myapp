package controller

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "myapp/api/v1"
	"myapp/internal/consts"
	"myapp/internal/dao"
	"myapp/internal/model/do"
	"myapp/internal/model/entity"
)

var Test = cTest{}

type cTest struct{}

func (c *cTest) Add(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	//获取token
	strId := gconv.String(ctx.Value(consts.CtxAdminId))
	get, _ := g.Redis().Get(ctx, "gfToken:"+strId)

	//修改token
	anyMap := make(map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(get.String()), &anyMap)
	g.Log().Info(ctx, "获取token的值：", anyMap)
	var user *entity.UserInfo
	_ = dao.UserInfo.Ctx(ctx).Where(do.UserInfo{Id: ctx.Value(consts.CtxAdminId)}).Scan(&user)
	anyMap["data"] = user
	g.Log().Info(ctx, "修改之后token的值：", anyMap)

	//修改之后的token写入redis
	_, _ = g.Redis().Set(ctx, "gfToken:"+strId, anyMap)

	return nil, err
}
