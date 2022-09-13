package main

import (
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	sql = "CREATE TABLE IF NOT EXISTS \"tbl_data\" (\n  \"id\" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,\n  \"ct\" integer,\n  \"data\" text\n);"
)

func InitDB() gdb.DB {
	db := g.DB()
	_, err := db.Exec(gctx.New(), sql)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := InitDB()
	print(db)
}
