package weapon

import (
	"game/info/tables"
	"math/rand"
	"time"
)

var (
	WuQie = tables.Weapon{
		Rank:        5,
		Name:        "雾切之回光",
		WeaponId:    1,
		RefineRange: 1,
		Grade:       1,
	}
	SongLai = tables.Weapon{
		Rank:        5,
		Name:        "松籁响起之时",
		WeaponId:    2,
		RefineRange: 1,
		Grade:       1,
	}
	XiaLiLongYin = tables.Weapon{
		Rank:        4,
		Name:        "匣里龙吟",
		WeaponId:    3,
		RefineRange: 1,
		Grade:       1,
	}
	JiLiCam = tables.Weapon{
		Rank:        4,
		Name:        "祭礼剑",
		WeaponId:    1,
		RefineRange: 1,
		Grade:       1,
	}
	YiLiFuRen = tables.Weapon{
		Rank:        3,
		Name:        "以理服人",
		WeaponId:    1,
		RefineRange: 1,
		Grade:       1,
	}
	LiMingShenCam = tables.Weapon{
		Rank:        3,
		Name:        "黎明神剑",
		WeaponId:    1,
		RefineRange: 1,
		Grade:       1,
	}
)

func ListReturn(rank int) tables.Weapon {
	rand.Seed(time.Now().Unix())
	var slice5 []tables.Weapon
	slice5 = append(slice5, WuQie, SongLai)
	var slice4 []tables.Weapon
	slice4 = append(slice4, XiaLiLongYin, JiLiCam)
	var slice3 []tables.Weapon
	slice3 = append(slice3, YiLiFuRen, LiMingShenCam)
	if rank == 4 {
		return slice4[rand.Intn(len(slice4))]
	}
	if rank == 3 {
		return slice3[rand.Intn(len(slice3))]
	}
	if rank == 5 {
		return slice5[rand.Intn(len(slice5))]
	}
	return tables.Weapon{}
}
