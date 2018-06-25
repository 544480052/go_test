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

/**
申明函数 【函数所属结构体】 函数名 【返回值类型】
 */
func (person Person) close()  {
	fmt.Println("I am close")
}
func (person Person) myAge() int {
	return person.age;
}

type Me struct {
	Person
	sex string
}





func main() {

	//给结构体复制
	//p:=Person{
	//	"张三",
	//	12,
	//	[]string{"aaa","bbb"},
	//}
	//
	//fmt.Println(p)

	//调用结构体函数
	//var person  = new(Person)
	//person.close()

	//var person = new(Person);
	//person.age = 12;
	//var age = person.myAge();
	//println("my age is ",age);//只能打印字符串和数字类型

	//结构体继承
	var me = new(Me)
	me.name = "cx"
	me.age = 18
	me.like = []string{"cccc","dddd"}
	fmt.Println("I like ",me.like)
	println("my age is ",me.age)
	fmt.Println("I am ",*me)

}
