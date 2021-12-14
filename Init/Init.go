package Init

import (
	"game/info/tables"
	"game/sqls"
)

func CreateInit() {
	alice := &tables.User{
		Password: "123456",
		Username: "alice",
		Uid:      "111",
		Account:  "123@qq.com",
	}
	bob := &tables.User{
		Password: "ABCDEF",
		Username: "bob",
		Uid:      "222",
		Account:  "456@qq.com",
	}
	cindy := &tables.User{
		Password: "abcdef",
		Username: "cindy",
		Uid:      "333",
		Account:  "789@qq.com",
	}
	sqldb := sqls.SQLDB
	sqldb.Db.Create(&alice)
	sqldb.Db.Create(&bob)
	sqldb.Db.Create(&cindy)
}
