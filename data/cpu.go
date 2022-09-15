package data

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"strconv"
	"time"
)

type CPUInfo struct {
	ModelName   string
	Count       int
	Temperature float32
	Usage       float64
}

func GetCPUInfo() *CPUInfo {
	cpuInfo := &CPUInfo{}
	counts, _ := cpu.Counts(false)
	percent, _ := cpu.Percent(time.Second, false)
	info, _ := cpu.Info()
	temperatures, _ := host.SensorsTemperatures()
	print(temperatures)
	cpuInfo.ModelName = info[0].ModelName
	cpuInfo.Count = counts
	cpuInfo.Usage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", percent[0]), 64)
	return cpuInfo
}
