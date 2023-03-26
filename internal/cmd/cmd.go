package cmd

import (
	"context"
	"github.com/gogf/gf/v2/net/goai"
	"myapp/internal/consts"
	"myapp/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gToken := NewGToken(ctx)
			s := g.Server()
			//设置全局的统一响应体中间件
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS)
				err := gToken.Middleware(ctx, group)
				if err != nil {
					g.Log().Errorf(ctx, "gToken中间件绑定group失败:%v", err)
					panic(err)
				}
				group.Bind(
					controller.User,
					controller.Test,
				)
			})

			//自定义增强API文档。
			enhanceOpenAPIDoc(s)
			//运行服务
			s.Run()
			return nil
		},
	}
)

/*
自定义增强API文档
*/
func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()

	//这里需要把中间件丢进来，否则统一的返回体不能展示
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = "Data"

	// 描述信息
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name:  "联系方式",
			Email: "jinhu007@outlook.com",
			URL:   "http://106.55.186.222",
		},
	}
}
