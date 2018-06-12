package main

import "fmt"

/**
	数组操作
	声明一个固定长度数组(srting是数组内元素类型，也可以是int或申明其他类型)
	var arr [3]string
	声明一个自适应长度数组
	var arr [...]string
	声明并赋值一个自适应长度数组
	arr := [...]int{1,2,3}

	切片操作
	声明切片
	var slice []int
	声明并赋值一个切片
	slice := []int{}	slice :=make([]int,5)

 */

func modify1(arr [5]int) {
	arr[0] = 10;
	fmt.Println("aaaaa1", arr)
}

func modify2(arr []int) {
	arr[0] = 10;
	fmt.Println("bbbbb1", arr)
}

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5} //数组
	modify1(arr1)                 //值传递
	fmt.Println("aaaaa2", arr1)
	arr2 := []int{1, 2, 3, 4, 5} //切片
	modify2(arr2)                //引用传递
	fmt.Println("bbbbb2", arr2)

}

func main() {
	//向数组中添加新元素
	arr2 := []int{1, 2, 3, 4, 5}
	var arr1 = arr2
	arr2 = append(arr2, 10)
	fmt.Println("aaaaa", arr1)
	fmt.Println("bbbbb", arr2)
}
