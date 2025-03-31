package main

import (
	_ "gfs/internal/logic"
	_ "gfs/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gfs/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
