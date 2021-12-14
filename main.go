package main

import (
	"game/Init"
	"game/handler/Get"
	"game/handler/login"
	"game/handler/selectcard"
	"game/sqls"

	"github.com/gin-gonic/gin"
)

func main() {
	/*因只有登录接口，故数据库中已提前预设三个默认账号
	alice:123456,Uid:111,account:123@qq.com
	bob:ABCDEF,Uid:222,account:456@qq.com
	cindy:abcdef,Uid:333,account:789@qq.com
	*/
	//初始化

	sqls.SetTheDeafult()
	Init.CreateInit()
	//设置路由

	r := gin.Default()
	passroute := r.Group("/pass")
	cardroute := r.Group("/card")
	getroute := r.Group("/get")
	login.Login(passroute)
	selectcard.Select(cardroute)
	Get.Rough(getroute)
	Get.GetDetails(getroute)
	r.Run()
	//for a part of test

	/*

		sqldb := sqls.InitConnection()
		alice := get.GetUser("123@qq.com")
		alice.Mora = 200
		sqldb.Db.Where("username = ?", "alice").Save(alice)*/
}
