package sql

const (
	CreateTableSql = "CREATE TABLE IF NOT EXISTS\"tbl_data\" (\n  \"id\" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,\n" +
		"  \"create_time\" integer,\n  \"data\" text,\n  \"data_type\" integer,\n  \"identify\" TEXT\n)"
)
