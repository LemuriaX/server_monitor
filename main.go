package main

import (
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"server_monitor/util"
)

func main() {
	db := util.InitDB()
	print(db)
}
