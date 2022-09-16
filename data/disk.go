package data

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
)

type DiskInfo struct {
	Name          string
	Temperature   float32
	Size          float32
	Usage         float32
	Free          float32
	OccupancyRate float32
	Read          float32
	Write         float32
}

func GetAllDisk() []string {
	resultString, _ := gproc.ShellExec(gctx.New(), "lsblk --j -d")
	resultJsons := gjson.New(resultString).Get("blockdevices").Array()
	var diskList []string
	for _, resultJson := range resultJsons {
		result := resultJson.(map[string]interface{})
		diskList = append(diskList, "/dev/"+result["name"].(string))
	}
	return diskList
}

func GetDiskInfo(diskName string) *DiskInfo {
	diskInfo := &DiskInfo{}
	diskInfo.Name = diskName

	return diskInfo
}
