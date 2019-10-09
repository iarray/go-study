package main

import "fmt"

type MyStruct struct{
	a int 
	b string
	c float32
}

func main(){
	i := 5
	p := &i

	*p = 6

	fmt.Println(p, i)

	var p2 *int = p

	*p2 = 8

	fmt.Println(p2, i)

	p3 := &p

	fmt.Println(p3, *p3, **p3)
	fmt.Printf("%T\n", p3)
	

	//error
	// var p4 *int = nil
	// *p4 = 0
	// fmt.Println(p4, *p4)

	str1 := MyStruct{1,"2",3.0}
	//引用复制还是值复制 ?
	str2 := str1
	str2.a = 4
	str2.b = "5"
	str2.c = 6.0

	//结果是值复制, str1和str2指向的内存是不同的
	//将str1赋给str2会两步走: 1.为str2分配空间, 2.把str1的值复制到str2
	//同样的, 返回结构体也会造成值复制, 
	fmt.Println(str1)
	fmt.Println(str2)
}
