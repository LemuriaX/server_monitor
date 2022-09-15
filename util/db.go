package util

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"server_monitor/sql"
)

func InitDB() gdb.DB {
	db := g.DB()
	_, err := db.Exec(gctx.New(), sql.CreateTableSql)
	if err != nil {
		panic(err)
	}
	return db
}
