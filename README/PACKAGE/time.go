package main

import (
	"fmt"
	"time"
)

func main() {

	//毫秒级时间戳
	fmt.Println(time.Now().UnixNano())

	//秒级时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//yyyy-mm-dd时间格式化成时间戳
	tm1, _ := time.Parse("2006-01-02 03:04:05", "2018-01-01 01:01:01")
	fmt.Println(tm1.Unix())

	//时间戳格式化成yyyy-mm-d
	//【时间的输出格式layout是有规定的：
	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday】
	tm2 := time.Unix(timestamp, 0)
	fmt.Println(tm2.Format("2006-01-02 03:04:05"))

}
