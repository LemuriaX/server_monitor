package data

type DiskInfo struct {
	Temperature   float32
	Size          float32
	Usage         float32
	Free          float32
	OccupancyRate float32
	Read          float32
	Write         float32
}

func GetDiskInfo() {
}
