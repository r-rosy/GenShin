package save

import (
	"game/sqls"
)

//must be a pointer of a struct
func Save(object interface{}, account string) {
	sqldb := sqls.SQLDB
	sqldb.Db.Where("account = ?", account).Save(object)
}
