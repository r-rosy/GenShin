package commonpass

func CheckJsonExist(json map[string]interface{}, head string) (interface{}, bool) {
	val, ok := json[head]
	return val, ok
}
