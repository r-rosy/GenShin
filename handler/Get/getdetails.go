package Get

import (
	"fmt"
	"game/handler/commonpass"
	"game/info/details"
	"game/info/tables"
	"game/special"

	"github.com/gin-gonic/gin"
)

func GetDetails(details *gin.RouterGroup) {
	json := make(map[string]interface{})
	handler := func(c *gin.Context) {
		account, bools := commonpass.CheckCookie(c)
		if bools {
			c.BindJSON(&json)
			val, ok := commonpass.CheckJsonExist(json, "type")
			if ok {
				Type, _ := val.(string)
				if Type != "character" && Type != "weapon" && Type != "godleftvalue" {
					c.JSON(400, gin.H{"message": "请输入正确类型"})
				} else {
					TypeSelect(account, c, Type)
				}
			} else {
				c.JSON(400, gin.H{"message": "some errors happend"})
			}
		} else {
			c.JSON(400, gin.H{"message": "some errors happend"})
		}
	}
	details.POST("/details", handler)
}
func TypeSelect(account string, c *gin.Context, Type string) {
	det := details.InqueryUpdate(account)
	var result = make(map[string]interface{})
	switch Type {
	case "character":
		character, _ := det["character"].([]tables.Character)
		for i, v := range character {
			m := special.StructToMap(v)
			stri := fmt.Sprintf("%d", i)
			result[stri] = m
		}
	case "weapon":
		weapon, _ := det["weapon"].([]tables.Weapon)
		for i, v := range weapon {
			m := special.StructToMap(v)
			stri := fmt.Sprintf("%d", i)
			result[stri] = m
		}
	case "godleftvalue":
		godleftvalue, _ := det["godleftvalue"].([]tables.GodLeftValue)
		for i, v := range godleftvalue {
			m := special.StructToMap(v)
			stri := fmt.Sprintf("%d", i)
			result[stri] = m
		}

	}
	c.JSON(200, result)
}
