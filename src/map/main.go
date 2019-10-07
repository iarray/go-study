package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	//1. make构建
	//注意: 不要使用 new，永远用 make 来构造 map
	map1 = make(map[string]int) //相当于 map[string]int{}
	//2. 初始化赋值
	map1 = map[string]int{"key2": 2, "key3": 3}
	map1["key1"] = 1
	fmt.Println(map1)

	//错误示范
	/*
		map2 := new(map[string]int)
		fmt.Printf("typeof(map2)=%T, &map2=%p \n", map2, map2)
		//会运行错误, map2是空指针
		(*map2)["k2"] = 2
		//报错panic: assignment to entry in nil map
		fmt.Println(map2)
		//注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	*/

	//判断某个 key 是否存在而不关心它对应的值到底是多少
	if value, isContainKey := map1["key1"]; isContainKey {
		fmt.Printf("key1存在, value=%d\n", value)
	}

	//删除键
	delete(map1, "key1")
	fmt.Println(map1)

	//遍历map
	for key, value := range map1 {
		//注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
		fmt.Printf("key is: %s, value is %d\n", key, value)
	}

	//map类型切片
	//初始化map切片需要调用一次make
	items := make([]map[int]int, 5)
	for i := range items {
		//元素初始化也需要调用一次make
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Value of items: %v\n", items)
}
