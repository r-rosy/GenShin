package selectcard

import (
	"game/handler/commonpass"

	"github.com/gin-gonic/gin"
)

func EstimateType(c *gin.Context, json map[string]interface{}, account string) string {
	val, ok := commonpass.CheckJsonExist(json, "type")
	if ok {
		v, _ := val.(string)
		switch v {
		case "character":
			return v
		case "weapon":
			return v
		case "godleftvalue":
			return v
		default:
			c.JSON(400, gin.H{"message": "some errors happend"})
			return ""
		}
	} else {
		c.JSON(400, gin.H{"message": "some errors happend"})
		return ""
	}
}
