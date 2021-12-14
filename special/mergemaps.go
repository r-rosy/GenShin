package special

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	Maps := make(map[string]interface{})
	for _, Map := range maps {
		for key, val := range Map {
			Maps[key] = val
		}
	}
	return Maps
}
