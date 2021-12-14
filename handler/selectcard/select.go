package selectcard

import (
	"fmt"
	"game/handler/commonpass"
	"game/info/cardpool"
	"game/info/details"
	"game/info/tables"
	"game/sqls"

	"github.com/gin-gonic/gin"
)

func Select(card *gin.RouterGroup) {
	json := make(map[string]interface{})
	handler := func(c *gin.Context) {
		account, bools := commonpass.CheckCookie(c)
		if bools {
			c.BindJSON(&json)
			val, ok := commonpass.CheckJsonExist(json, "number")
			if ok {
				num, _ := val.(float64)
				number := int(num)
				if number != 1 && number != 10 {
					c.JSON(400, gin.H{"message": "只能一次抽一连或十连"})
				} else {
					TypeSelect(account, c, number, json)
				}
			} else {
				c.JSON(400, gin.H{"message": "some errors happend"})
			}
		} else {
			c.JSON(400, gin.H{"message": "some errors happend"})
		}
	}
	card.POST("/select", handler)
}
func Failed(c *gin.Context, num int) {
	sentense := fmt.Sprintf("原石不足，请先获取原石,您的剩余原石是%d", num)
	c.JSON(200, gin.H{"message": sentense})
}
func TypeSelect(account string, c *gin.Context, num int, json map[string]interface{}) {
	det := details.InqueryUpdate(account)
	user, _ := det["user"].(*tables.User)
	if user.Rough < ((160) * num) {
		Failed(c, user.Rough)
	} else {
		Type := EstimateType(c, json, account)
		if Type == "" {

		} else {
			rank := cardpool.RankRadom(user)
			//目前只能给我单抽
			if num == 1 {
				user.Rough = user.Rough - 160
			} /*else {
				user.Rough=user.Rough-1600
			}*/
			sqls.SQLDB.Db.Where("account = ?", account).Save(user)
			Map := cardpool.CSelectResults(num, rank, account, Type)
			c.JSON(200, Map)
		}
	}
}
