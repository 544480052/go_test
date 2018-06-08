package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	//for i:=0;i<10;i++ {
	//	//生成0-100之间的整数随机数
	//	fmt.Println(rand.Intn(100))
	//}

	//初始时间化随机数种子
	//rand.Seed(time.Now().UnixNano())
	//for i:=0;i<10;i++  {
	//	fmt.Println(rand.Intn(100))
	//}

	//初始时间化随机数种子
	//r:=	rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i:=0;i<10;i++  {
	//	fmt.Println(r.Intn(100))
	//}


	for i:=0;i<10;i++  {
		fmt.Println(rand.Int63())
	}
	//初始化时间随机数种子
	r:=	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<10;i++  {
		fmt.Println(r.Int63())
	}





}

