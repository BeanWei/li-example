package main

import (
	"github.com/BeanWei/li-example/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Li.Run(gctx.New())
}
