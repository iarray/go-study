package main

import (
	"fmt"
	"time"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	var c1 chan string
	c1 = make(chan string)       //只能传输字符串的通道
	c2 := make(chan interface{}) //可以传输任意数据的通道

	//相当于c#的Task.Run
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		c1 <- "hello world"
	}()

	//这里会阻塞, 直到接收到c1的值
	//相当于c#的 await Task<string>
	dataC1 := <-c1
	fmt.Println("Recieve c1=", dataC1)

	//传递任意类型
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		c2 <- &Dog{Name: "旺旺", Age: 6}
		time.Sleep(time.Duration(1) * time.Second)
		c2 <- "fuck"
	}()

	fmt.Println("Recieve c2=", <-c2)
	fmt.Println("Recieve c2=", <-c2)

	//遍历管道
	cn := make(chan int)

	go func() {

		for i := 3; i > 0; i-- {
			cn <- i
			time.Sleep(time.Duration(1) * time.Second)
		}

	}()

	//如果goroutines执行完成还在等待管道数据接收则会报错
	//fatal error: all goroutines are asleep - deadlock!
	for c := range cn {
		fmt.Println(c)

		if c == 0 {
			break
		}
	}
}
