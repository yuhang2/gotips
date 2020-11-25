package main

import (
	"github.com/yuhang2/gotips/datarace/grabx"
)

func main() {
	grabx.Initialize()

	grabx.SyncGetX()
}
