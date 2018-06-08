package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {

	//for i:=0;i<10;i++ {
	//	//生成0-100之间的整数随机数
	//	fmt.Println(rand.Intn(100))
	//}

	for i:=0;i<10;i++  {
		fmt.Print(rand.Seed(time.Now().UnixNano()))
	}

}



