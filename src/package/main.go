package main

import(
	"../github.com/iarray/hello"
	"fmt"
)

var a string = "G"

func init()  {
	fmt.Println("go自动调用了init， 并且在main之前执行")
}

func main()  {
	fmt.Printf("PI=%f\n", hello.PI)
	fmt.Println(hello.GetName())
	n()
	m()
	n()
}


func n() { print(a) }

func m() {
   a := "O"
   print(a)
}