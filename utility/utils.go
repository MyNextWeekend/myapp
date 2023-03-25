package utility

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

func CfgGet(ctx context.Context, key string) *gvar.Var {
	value, err := g.Cfg().Get(ctx, key)
	if err != nil {
		g.Log().Info(ctx, "从配置文件中读取【%v】信息错误", key)
		panic(err)
	}
	return value
}
