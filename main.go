package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	data "server_monitor/data"
	"server_monitor/util"
	"time"
)

func main() {
	ctx := gctx.New()
	_ = util.InitDB()
	cpuInfo := data.GetCPUInfo()
	fmt.Printf("%v\n", cpuInfo)
	diskMap := data.GetAllDisk()
	gcache.Set(ctx, "diskMap", diskMap, time.Minute)
	for key, _ := range diskMap {
		diskInfo := data.GetDiskInfo(ctx, key)
		fmt.Printf("%v\n", diskInfo)
	}

}
