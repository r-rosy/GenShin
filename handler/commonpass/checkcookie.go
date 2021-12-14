package commonpass

import (
	"game/info/tables"
	"game/sqls"

	"github.com/gin-gonic/gin"
)

func CheckCookie(c *gin.Context) (string, bool) {
	cookie := &tables.Cookie{}
	cook, err := c.Request.Cookie("OriginGod2021")
	val := cook.Value
	if err != nil {
		return "", false
	}
	sqldb := sqls.SQLDB
	if err := sqldb.Db.Where("value = ?", val).Find(cookie).Error; err != nil {
		return "", false
	} else {
		return cookie.Account, true
	}
}
