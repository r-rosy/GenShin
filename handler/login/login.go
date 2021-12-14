package login

import (
	"game/info/tables"
	"game/special"
	"game/sqls"

	"github.com/gin-gonic/gin"
)

func Login(passroute *gin.RouterGroup) {
	sqldb := sqls.SQLDB
	json := make(map[string]string)
	handler := func(c *gin.Context) {
		c.BindJSON(&json)
		val1, ok1 := CheckJsonExist(json, "account")
		val2, ok2 := CheckJsonExist(json, "password")
		if ok1 && ok2 {
			if val1 == "" || val2 == "" {
				c.JSON(200, gin.H{"message": "请同时输入账号及密码"})
			} else {
				if CheckCorrect(json) {
					s := special.RandStringBytesMaskImprSrcUnsafe(20)
					cookie := &tables.Cookie{
						Name:    "OriginGod2021",
						Value:   s,
						Account: val1,
					}
					cook := &tables.Cookie{}
					for sqldb.Db.Where("account = ?", val1).Find(cook); cook.Account != ""; {
						sqldb.Db.Where("account = ?", val1).Delete(cook)
						cook = &tables.Cookie{}
					}
					sqldb.Db.Create(cookie)
					c.SetCookie("OriginGod2021", s, 3600, "/", "localhost", false, true)
					c.JSON(200, gin.H{"message": "登录成功"})
				} else {
					c.JSON(200, gin.H{"message": "密码错误，登录失败"})
				}
			}
		} else {
			c.JSON(400, gin.H{"message": "some errors happend"})
		}

	}
	passroute.POST("/login", handler)
}
