package data

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/net"
)

type NetworkInfo struct {
	Name   string
	Input  float32
	Output float32
}

func GetAllNetworkInterface() []string {
	interfaces, _ := net.Interfaces()
	var networkInterfaces []string
	for _, networkInterface := range interfaces {
		networkInterfaces = append(networkInterfaces, networkInterface.Name)
	}
	return networkInterfaces
}

func GetNetworkInfo() []NetworkInfo {
	var networkInfoList []NetworkInfo
	res, _ := gproc.ShellExec(gctx.New(), "sar -n DEV 1 1 | grep \"Average\"")
	split := gstr.Split(res, "\n")
	for _, info := range split {
		if gstr.LenRune(info) == 0 {
			continue
		}
		titleStr := gstr.Replace(info, "Average:", "")
		data := gstr.SplitAndTrim(titleStr, " ")
		if gstr.IsNumeric(data[1]) {
			networkInfo := NetworkInfo{}
			networkInfo.Name = data[0]
			networkInfo.Input = gconv.Float32(data[3])
			networkInfo.Output = gconv.Float32(data[4])
			networkInfoList = append(networkInfoList, networkInfo)
		}
	}
	return networkInfoList
}
