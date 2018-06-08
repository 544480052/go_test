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

	//r:=	rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++  {
		fmt.Println(rand.Intn(100))
	}

}

