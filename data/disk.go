package data

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"strconv"
)

type DiskInfo struct {
	Name        string
	Temperature int
	Size        string
	// Usage         float32
	// Free          float32
	// OccupancyRate float32
	Read  float32
	Write float32
}

func GetAllDisk() map[string]float64 {
	resultString, _ := gproc.ShellExec(gctx.New(), "lsblk --j -db")
	resultJsons := gjson.New(resultString).Get("blockdevices").Array()
	diskMap := map[string]float64{}
	for _, resultJson := range resultJsons {
		result := resultJson.(map[string]interface{})
		diskMap["/dev/"+result["name"].(string)] = result["size"].(float64)
	}
	return diskMap
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

func GetDiskInfo(ctx context.Context, diskName string) *DiskInfo {
	diskInfo := &DiskInfo{}
	diskInfo.Name = diskName
	diskInfo.Temperature = GetDiskTemperature(diskName)
	diskMap, _ := gcache.Get(ctx, "diskMap")
	size := diskMap.MapStrAny()[diskName].(float64) / (1024 * 1024)
	if size > 1024 {
		diskInfo.Size = fmt.Sprintf("%.2f GB", size/1024)
	} else {
		diskInfo.Size = fmt.Sprintf("%.2f MB", size)
	}

	return diskInfo
}
