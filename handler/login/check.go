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
	if err := sqldb.Db.Where("account = ?", json["account"]).First(&user).Error; err != nil {
		panic(err)
	} else {
		if json["password"] == user.Password {
			return true
		} else {
			return false
		}
	}
}
