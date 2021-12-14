package character

import (
	"game/info/tables"
	"math/rand"
	"time"
)

var (
	KeQing = tables.Character{
		Name:        "刻晴",
		Rank:        5,
		Grade:       1,
		RefineRange: 0,
	}
	ShenliLingHua = tables.Character{
		Name:        "神里凌华",
		Rank:        5,
		Grade:       1,
		RefineRange: 0,
	}
	BanNiTe = tables.Character{
		Name:        "班尼特",
		Rank:        4,
		Grade:       1,
		RefineRange: 0,
	}
	TuoMa = tables.Character{
		Name:        "托马",
		Rank:        4,
		Grade:       1,
		RefineRange: 0,
	}
)

func ListReturn(rank int) tables.Character {
	rand.Seed(time.Now().Unix())
	var slice5 []tables.Character
	slice5 = append(slice5, KeQing, ShenliLingHua)
	var slice4 []tables.Character
	slice4 = append(slice4, BanNiTe, TuoMa)
	if rank == 4 {
		return slice4[rand.Intn(len(slice4))]
	} else if rank == 5 {
		return slice5[rand.Intn(len(slice5))]
	} else if rank == 3 {
		return tables.Character{}
	}
	return tables.Character{}
}
