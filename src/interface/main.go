package main

import(
	"fmt"
)

type Man struct{
	Name string
}

func main()  {
	var i interface{} = 88
	b := 12
	//报错
	//fmt.Println(i + b)
	fmt.Println(i.(int) + b)

	var obj interface{} = Man{Name:"aa"}
	obj2 := obj.(Man) 
	obj2.Name = "aaaa" 
	fmt.Println(obj)
	fmt.Println(obj2)
}