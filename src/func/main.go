package main

import (
	"fmt"
	"io"
	"log"
)

type binOp func(int, int) int

func main(){
	var f binOp;
	f = add

	fmt.Println(f(1,2))

	arr := []int{1,2,3,4,5,6}
	m := min(arr...)
	fmt.Println(m)
	m = min(9,7,5,7,3,4)
	fmt.Println(m)
	
	typecheck(1, 3.14, float64(0.678), "kkkk", true, min)

	function1()

	func1("Go")
}

func add(a int, b int) int {
	return a+b
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

/* 空接口可以接受任何类型的参数 */
func typecheck(values ... interface{}) {
    for _, value := range values {
        switch value.(type) {
            case int: fmt.Println("is int")
            case float32, float64:fmt.Println("is float")
            case string: fmt.Println("is string")
            case bool: fmt.Println("is bool")
            default: fmt.Println("unknow")
        }
    }
}

/* defer相当于java的try finally的finally */
func function1() {
	fmt.Printf("In function1 at the top\n")
	//defer调用的函数会在当前函数完成后执行, 相当于函数入栈,遵循后进先出
	defer function2()
	defer function3()	

	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Println("Function2: Deferred until the end of the calling function!")
}

func function3() {
	fmt.Println("Function3: Deferred until the end of the calling function!")
}


//使用 defer 语句来记录函数的参数与返回值
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

/**
close 	用于管道通信
len、cap 	len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
new、make 	new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号
copy、append 	用于复制和连接切片
panic、recover 	两者均用于错误处理机制
print、println 	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
complex、real imag 	用于创建和操作复数（详见第 4.5.2.2 节）
*/