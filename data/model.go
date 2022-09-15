package data

type Info interface {
}

type TblDataModel struct {
	id         int
	createTime int
	data       Info
	dataType   int
	identify   string
}
