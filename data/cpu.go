package data

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"strconv"
	"time"
)

type CPUInfo struct {
	ModelName   string
	Count       int
	Temperature float64
	Usage       float64
}

func GetCPUInfo() *CPUInfo {
	cpuInfo := &CPUInfo{}
	counts, _ := cpu.Counts(false)
	percent, _ := cpu.Percent(time.Second, false)
	info, _ := cpu.Info()
	temperatures, _ := host.SensorsTemperatures()
	for _, temperature := range temperatures {
		// look at
		// https://www.kernel.org/doc/html/v5.12/hwmon/k10temp.html
		// https://www.kernel.org/doc/html/v5.12/hwmon/k8temp.html
		// https://www.kernel.org/doc/html/v5.12/hwmon/coretemp.html
		if gstr.Contains(temperature.SensorKey, "k10temp") ||
			gstr.Contains(temperature.SensorKey, "k8temp") ||
			gstr.Contains(temperature.SensorKey, "coretemp") {
			cpuInfo.Temperature, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temperature.Temperature), 64)
		}
	}
	cpuInfo.ModelName = info[0].ModelName
	cpuInfo.Count = counts
	cpuInfo.Usage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", percent[0]), 64)
	return cpuInfo
}
