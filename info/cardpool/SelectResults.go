package cardpool

import (
	"game/special"
	"game/sqls"
	"game/warehouse/character"
	"game/warehouse/godleftvalue"
	"game/warehouse/weapon"
)

func CSelectResults(num int, rank int, account string, Type string) map[string]interface{} {
	if num == 1 {
		return selectOne(rank, account, Type)
	} else {
		return map[string]interface{}{"message": "一次只能抽一次"}
	}
}
func selectOne(rank int, account string, Type string) map[string]interface{} {
	switch Type {
	case "character":
		if rank != 3 {
			chara := character.ListReturn(rank)
			chara.Account = account
			sqldb := sqls.SQLDB
			sqldb.Db.Create(&chara)
			maps := special.StructToMap(chara)
			return maps
		} else {
			return map[string]interface{}{"message": "三星武器，你懂的..."}
		}
	case "weapon":

		weapon := weapon.ListReturn(rank)
		weapon.Account = account
		sqldb := sqls.SQLDB
		sqldb.Db.Create(&weapon)
		maps := special.StructToMap(weapon)
		return maps
	case "godleftvalue":
		godleftvalue := godleftvalue.ListReturn(rank)
		godleftvalue.Account = account
		sqldb := sqls.SQLDB
		sqldb.Db.Create(&godleftvalue)
		maps := special.StructToMap(godleftvalue)
		return maps
	}
	return map[string]interface{}{}

}

/*
//失败的功能
func selectTen(rank int, account string, Type string) map[string]interface{} {
	Maps := make(map[string]interface{})
	for i := 0; i <= 9; i++ {
		Map := selectOne(rank, account, Type)
		Maps = special.MergeMaps(Map, Maps)
		fmt.Println(Maps)
	}
	return Maps
}*/
