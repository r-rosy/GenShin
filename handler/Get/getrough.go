package Get

import (
	"fmt"
	"game/handler/commonpass"
	"game/info/execute/get"
	"game/info/execute/save"

	"github.com/gin-gonic/gin"
)

func Rough(Get *gin.RouterGroup) {
	json := make(map[string]interface{})
	handler := func(c *gin.Context) {
		account, bools := commonpass.CheckCookie(c)
		if bools {
			c.BindJSON(&json)
			val, ok := commonpass.CheckJsonExist(json, "number")
			if ok {
				num, _ := val.(float64)
				number := int(num)
				user := get.GetUser(account)
				rough := user.Rough
				user.Rough = rough + number
				save.Save(user,account)
				sentense := fmt.Sprintf("充值创世结晶成功,您的剩余原石是%d",user.Rough)
				c.JSON(200, gin.H{"message": sentense})
			} else {
				c.JSON(400, gin.H{"message": "some errors happend"})
			}
		} else {
			c.JSON(400, gin.H{"message": "some errors happend"})
		}
	}
	Get.POST("/rough", handler)
}
