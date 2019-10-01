package main

import "fmt"

const(
	A = iota	//0	
	B = 50
	C			//50
	D			//50
	E			//50
	F = iota	//5
	G			//6
	H = iota	//7
	I			//8
	J			//9
)

type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type BitFlag int
const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // 1 << 1 == 2
	Receive // 1 << 2 == 4
)


func main()  {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)

	flag := Active | Send // == 3
	fmt.Println(Active)
	fmt.Println(Send)
	fmt.Println(Receive)
	fmt.Println(flag)

	a:=10>5
	fmt.Println(a)
}