package data

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type MemoryInfo struct {
	Size  int
	Usage int
	Free  int
}

func GetMemoryInfo() []map[string]float32 {
	resultString, _ := gproc.ShellExec(gctx.New(), "free | egrep \"Mem|Swap\"")
	split := gstr.SplitAndTrim(resultString, "\n")
	var res []map[string]float32
	for _, info := range split {
		if gstr.LenRune(info) == 0 {
			continue
		}
		temp := make(map[string]float32)
		name := gstr.SplitAndTrim(info, ":")[0]
		data := gstr.SplitAndTrim(gstr.SplitAndTrim(info, ":")[1], " ")
		temp["total"] = gconv.Float32(data[0])
		temp["used"] = gconv.Float32(data[1])
		temp["free"] = gconv.Float32(data[2])
		if name == "Mem" {
			temp["shared"] = gconv.Float32(data[3])
			temp["buff/cache"] = gconv.Float32(data[4])
			temp["available"] = gconv.Float32(data[5])
		}
		res = append(res, temp)
	}
	return res
}
