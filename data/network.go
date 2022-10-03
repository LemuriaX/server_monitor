package data

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/net"
)

type NetworkInfo struct {
	Name      string
	Input     float32
	Output    float32
	SumInput  float64
	SumOutput float64
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
	sumRes, _ := gproc.ShellExec(gctx.New(), "cat /proc/net/dev | grep \":\"")
	split := gstr.Split(res, "\n")
	sumResSplit := gstr.Split(sumRes, "\n")
	networkMap := make(map[string][]float64)
	for _, info := range sumResSplit {
		if gstr.LenRune(info) == 0 {
			continue
		}
		name := gstr.SplitAndTrim(gstr.Split(info, ":")[0], " ")[0]
		var data []float64
		input := gstr.SplitAndTrim(gstr.Split(info, ":")[1], " ")[0]
		data = append(data, gconv.Float64(input))
		output := gstr.SplitAndTrim(gstr.Split(info, ":")[1], " ")[8]
		data = append(data, gconv.Float64(output))
		networkMap[name] = data
	}
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
			networkInfo.SumInput = networkMap[data[0]][0]
			networkInfo.SumOutput = networkMap[data[0]][1]
			networkInfoList = append(networkInfoList, networkInfo)
		}
	}
	return networkInfoList
}
