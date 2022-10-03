package data

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

type DiskInfo struct {
	Name        string
	Temperature int
	// Size        string
	// Usage         float32
	// Free          float32
	// OccupancyRate float32
	// Read  float32
	// Write float32
}

func GetDiskSize() map[string]map[string]float64 {
	resultString, _ := gproc.ShellExec(gctx.New(), "df | grep \"^/dev/\"")
	split := gstr.Split(resultString, "\n")
	res := make(map[string]map[string]float64)
	for _, info := range split {
		if gstr.LenRune(info) == 0 {
			continue
		}
		diskInfo := gstr.SplitAndTrim(info, " ")

		if diskInfo[5] == "/boot/efi" || diskInfo[5] == "/etc/pve" {
			continue
		}
		r := make(map[string]float64)
		r["used"] = gconv.Float64(diskInfo[2])
		r["Available"] = gconv.Float64(diskInfo[3])
		res[diskInfo[0]] = r
	}
	return res
}

func GetAllDisk() []string {
	resultString, _ := gproc.ShellExec(gctx.New(), "lsblk --j -db")
	resultJsons := gjson.New(resultString).Get("blockdevices").Array()
	var diskList []string
	for _, resultJson := range resultJsons {
		result := resultJson.(map[string]interface{})
		diskList = append(diskList, "/dev/"+result["name"].(string))
	}
	return diskList
}

func GetDiskTemperature(diskName string) int {
	result := 0
	if gstr.Contains(diskName, "nvme") {
		resultString, _ := gproc.ShellExec(gctx.New(), "nvme smart-log "+diskName+" | grep '^temperature' ")
		result, _ = strconv.Atoi(gstr.Trim(gstr.StrTillEx(gstr.StrEx(resultString, ":"), "C")))

	} else {
		resultString, _ := gproc.ShellExec(gctx.New(), "hddtemp -n "+diskName)
		resultString = gstr.Trim(resultString, "\n")
		if gstr.IsNumeric(resultString) {
			result, _ = strconv.Atoi(resultString)
		}
	}

	return result
}

func GetDiskInfo(diskName string) *DiskInfo {
	diskInfo := &DiskInfo{}
	diskInfo.Name = diskName
	diskInfo.Temperature = GetDiskTemperature(diskName)
	return diskInfo
}
