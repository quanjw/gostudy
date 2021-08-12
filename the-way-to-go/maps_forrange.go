package main

import (
	"fmt"
	"github.com/issue9/cnregion"
	"log"
)

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}


	v, err := cnregion.LoadFile("./data/regions.db", "-", 2020)

	if  err!=nil {
		log.Printf("Hash write error: %v / ", err)
	}

	p := v.Provinces() // 返回所有省列表
	cities := p[0].Items() // 返回该省下的所有市
	counties := cities[0].Items() // 返回该市下的所有县
	towns := counties[0].Items() // 返回所有镇
	villages := towns[0].Items() // 所有村和街道信息

	log.Print(villages)

	d := v.Districts() // 按以前的行政大区进行划分
	provinces := d[0].Items() // 该大区下的所有省份
	log.Print(provinces)

	list := v.Search("温州", nil) // 按索地名中带温州的区域列表
	log.Print(list)
}