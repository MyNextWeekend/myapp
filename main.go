package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "myapp/internal/logic"
	_ "myapp/internal/packed"

	"myapp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
