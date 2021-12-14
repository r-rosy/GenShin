package special

import (
	"encoding/json"
)

func StructToMap(Struct interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(Struct)
	json.Unmarshal(j, &m)
	return m
}
