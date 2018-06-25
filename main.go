package main

import "fmt"

/**
	&符号的意思是对变量取地址/指针，如：变量a的地址/指针是&a
	*符号的意思是对`指针`取值，如:*&a，就是a变量所在地址的值，当然也就是a的值
	*&可以抵消掉，但&*是不可以抵消
 */

type Person struct {
	name string
	age int
	like []string
}

func (person Person) close()  {
	fmt.Println("I am close")
}



func main() {

	//p:=Person{
	//	"张三",
	//	12,
	//	[]string{"aaa","bbb"},
	//}
	//
	//fmt.Println(p)

	var person  = new(Person)
	person.close()

}
