package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"server_monitor/data"
	"server_monitor/util"
)

func main() {
	_ = util.InitDB()
	cpuInfo := data.GetCPUInfo()
	fmt.Printf("%v\n", cpuInfo)
	diskInfo := data.GetDiskInfo()
	fmt.Printf("%v\n", diskInfo)
}
