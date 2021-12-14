package details

import (
	"game/info/tables"
	"game/sqls"
)

func InqueryUpdate(account string) map[string]interface{} {
	sqldb := sqls.SQLDB
	var user = &tables.User{}
	var cookie []tables.Cookie
	var weapon []tables.Weapon
	var godleftvalue []tables.GodLeftValue
	var character []tables.Character
	sqldb.Db.Where("account = ?", account).First(user)
	sqldb.Db.Where("account = ?", account).Find(&cookie)
	sqldb.Db.Where("account = ?", account).Find(&weapon)
	sqldb.Db.Where("account = ?", account).Find(&godleftvalue)
	sqldb.Db.Where("account = ?", account).Find(&character)
	details := make(map[string]interface{})
	details["user"] = user
	details["cookie"] = cookie
	details["weapon"] = weapon
	details["godleftvalue"] = godleftvalue
	details["character"] = character
	return details
}
