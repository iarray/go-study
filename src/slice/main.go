package main

import (
	"fmt"
)

func main(){
	a := []int{1,2,3}
	b := []int{4,5,6}

	//向切片添加元素
	c := append(a, 4, 5, 6)
	fmt.Println(c)
	fmt.Printf("&a=%p, &c=%p\n", &a, &c) //超过切片容量, 所以新建了一个新切片存放

	//如果没超切片容量,则在原切片上添加元素
	a2 := make([]int, 0, 10)
	tmp := append(a2[:], 1, 2, 3, 4, 5)
	fmt.Println(a2)
	fmt.Printf("&a2=%p, &tmp=%p\n", &a2, &tmp)

	

	//a和b合并成一个新切片
	d := append(a,b...)
	fmt.Println(d)

	//复制
	arr_from := []int{1,2,3}
	arr_to := make([]int ,10)

	copy(arr_to, arr_from)
	fmt.Println(arr_to)

}