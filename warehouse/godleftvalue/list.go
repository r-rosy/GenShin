package godleftvalue

import (
	"game/info/tables"
	"math/rand"
	"time"
)
var (
	MoNvFlower= tables.GodLeftValue{
		Grade: 1, 
		Rank:  5,
		Name:  "魔女花",
		Index: 1,
	}
	DuHuoCup =  tables.GodLeftValue{
		Grade: 1,
		Rank:  5,
		Name:  "渡火杯",
		Index: 3,
	}
	FightMadHead = tables.GodLeftValue{
		Grade: 1,
		Rank:   4,
		Name: "战狂头",
		Index: 4,
	} 
	YueTuanFlower = tables.GodLeftValue{
		Grade: 1,
		Rank: 4,
		Name:   "乐团花",
		Index: 1,
	}
	YouYiFeature = tables.GodLeftValue{
		Grade: 1,
		Rank:  3,
		Name:  "游医羽",
		Index: 2,
	}  
	YouYiCup =tables.GodLeftValue{
		Grade: 1, 
		Rank:  3,
		Name:  "游医杯",
		Index: 3,
	}
)	

func ListReturn(rank int) tables.GodLeftValue {
	rand.Seed(time.Now().Unix())
	var slice5 []tables.GodLeftValue
	slice5 = append(slice5, MoNvFlower,DuHuoCup)
	var slice4 []tables.GodLeftValue 
    slice4 = append(slice4, FightMadHead, YueTuanFlower)
	var slice3 []tables.GodLeftValue 
	slice3 = append(slice3, YouYiCup,YouYiFeature)
	if rank == 4 { 
		return slice4[rand.Intn(len(slice4))]
	}
	if rank == 3 {
		return slice3[rand.Intn(len(slice3))]
	}
	if rank == 5 {
		return slice5[rand.Intn(len(slice5))]
	}
	return tables.GodLeftValue{}
}
