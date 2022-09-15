package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"server_monitor/data"
	"server_monitor/util"
)

func main() {
	db := util.InitDB()
	print(db)
	cpuInfo := data.GetCPUInfo()
	fmt.Printf("%v", cpuInfo)
}
