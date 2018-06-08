/**
"math/rand"包 实现了伪随机数生成器。也就是生成 整形和浮点型
该包中根据生成伪随机数是是否有种子(可以理解为初始化伪随机数)，可以分为两类：
1、有种子。通常以时钟，输入输出等特殊节点作为参数，初始化。该类型生成的随机数相比无种子时重复概率较低。
2、无种子。可以理解为此时种子为1， Seed(1)
*/

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
