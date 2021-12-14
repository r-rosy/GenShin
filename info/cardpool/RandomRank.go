package cardpool

import (
	"game/info/tables"
)

func RankRadom(user *tables.User) int {
	num1 := user.Number1
	num2 := user.Number2
	user.Number1 = num1 + 1
	user.Number2 = num2 + 1
	if num2 < 90 {
		if num1 == 10 {
			clear(user, 4)
			return 4
		} else {
			rank := Random()
			if rank == 4 {
				clear(user, 4)
			}
			if rank == 5 {
				clear(user, 5)
			}
			return rank
		}
	} else {
		clear(user, 5)
		return 5
	}
}
func clear(user *tables.User, rank int) {
	if rank == 4 {
		user.Number1 = 1
	} else {
		user.Number1 = 1
		user.Number2 = 1
	}
}
