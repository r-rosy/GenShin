package get

import (
	"game/info/details"
	"game/info/tables"
)

func GetUser(account string) *tables.User {
	det := details.InqueryUpdate(account)
	val, _ := det["user"].(*tables.User)
	return val
}
func GetCharacter(account string) *tables.Character {
	det := details.InqueryUpdate(account)
	val, _ := det["character"].(*tables.Character)
	return val
}
func GetCookie(account string) *tables.Cookie {
	det := details.InqueryUpdate(account)
	val, _ := det["cookie"].(*tables.Cookie)
	return val
}
func GetWeapon(account string) *tables.Weapon {
	det := details.InqueryUpdate(account)
	val, _ := det["weapon"].(*tables.Weapon)
	return val
}
func GetGodLeftValue(account string) *tables.GodLeftValue {
	det := details.InqueryUpdate(account)
	val, _ := det["godleftvalue"].(*tables.GodLeftValue)
	return val
}
