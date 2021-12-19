package login

import (
	"game/info/tables"
	"game/sqls"
)

func CheckJsonExist(json map[string]string, head string) (string, bool) {
	val, ok := json[head]
	return val, ok
}
func CheckCorrect(json map[string]string) bool {
	user := &tables.User{}
	sqldb := sqls.SQLDB
	if sqldb.Db.Where("account = ?", json["account"]).First(&user); user.Account == "" {
		panic("error happend")
	} else {
		if json["password"] == user.Password {
			return true
		} else {
			return false
		}
	}
}
