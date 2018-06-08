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
	tm2 := time.Unix(timestamp, 0)
	fmt.Println(tm2.Format("2006-01-02 03:04:05"))

}
