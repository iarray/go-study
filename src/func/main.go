package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"runtime"
)

type binOp func(int, int) int


var(
	//定义一个全局函数, 用于打印代码位置
	where = func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
)

func main(){
	//自定义函数的类型
	var f binOp;
	f = add
	fmt.Println(f(1,2))

	//找出最小值
	arr := []int{1,2,3,4,5,6}
	m := min(arr...)
	fmt.Println(m)
	
	//定义可接受任意类型参数的函数
	typecheck(1, 3.14, float64(0.678), "kkkk", true, min)

	//defer 执行顺序测试
	function1()

	//defer技巧, 作为log输出
	myLog("Go")

	//找出数组内大于2的元素
	ret := indexFunc(arr, func (v int) bool  {
		return v>2
	})

	fmt.Println(ret)

	//defer 经常配合匿名函数使用, 它可以用于改变函数的命名返回值。
	fmt.Printf("tmpfunc()=%d\n", tmpfunc())

	//将函数作为返回值
	fmt.Printf("getAddFunc()(1,2)=%d\n", getAddFunc()(1,2))

	//通过函数内声明局部变量保存状态
	sumFunc := summer()
	fmt.Printf("sumFunc(1)=%d, 1=1\n", sumFunc(1))
	fmt.Printf("sumFunc(2)=%d, 1+2=3\n", sumFunc(2))
	fmt.Printf("sumFunc(3)=%d, 1+2+3=6\n", sumFunc(3))

	//通过工厂函数构建函数
	addJpg := MakeAddSuffix(".jpg")
	addPng := MakeAddSuffix(".png")
	fmt.Printf("addJpg(file)=%s\n", addJpg("file"))
	fmt.Printf("addPng(file)=%s\n", addPng("file"))
	
	//打印代码位置
	where()
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
func myLog(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

/* 函数作为参数 */
func indexFunc(s []int,f func(v int)bool) []int{
	res := make([]int, 0, len(s))
	for _,val := range s{
		fmt.Printf("val=%d, len(res)=%d, cap(res)=%d\n", val, len(res), cap(res))
		if f(val) {
			res = append(res, val)	
		}
	}
	return res
}

/* 匿名函数 */
func tmpfunc() (ret int) {
	defer func() {
		ret++
	}()	//别漏了最后的括号, 不然只是定义了,没有调用

	return 1 //实际返回2
} 

/* 将函数作为返回值 */
func getAddFunc() binOp{
	return func(x,y int)int{
		return x + y
	}
}

func summer() func(x int)int{
	var res int
	return func (x int)int{
		res += x
		return res
	}
}

/* 利用闭包原理定义工厂函数 */
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/******************		 go 内建函数(不需要引入包,自带函数)    **********

close 	用于管道通信
len、cap 	len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
new、make 	new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号
copy、append 	用于复制和连接切片
panic、recover 	两者均用于错误处理机制
print、println 	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
complex、real imag 	用于创建和操作复数（详见第 4.5.2.2 节）
************************		END			***********************/
