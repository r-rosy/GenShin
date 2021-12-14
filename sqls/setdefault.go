package sqls

var SQLDB *SqlDb

func SetTheDeafult() {
	sqldb := tableinit()
	SQLDB = sqldb
}
