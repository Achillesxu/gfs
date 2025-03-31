package main

import (
	_ "gfs/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfs/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
