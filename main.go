package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	data "server_monitor/data"
	"server_monitor/util"
)

func main() {
	_ = util.InitDB()
	cpuInfo := data.GetCPUInfo()
	fmt.Printf("%v\n", cpuInfo)
	diskList := data.GetAllDisk()
	for _, key := range diskList {
		diskInfo := data.GetDiskInfo(key)
		fmt.Printf("%v\n", diskInfo)
	}
	// res := data.GetAllNetworkInterface()
	networkInfo := data.GetNetworkInfo()
	fmt.Printf("%v", networkInfo)

}
