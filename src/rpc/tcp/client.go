package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	Val1, Val2 int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	var result int
	err = client.Call("MathMethod.Add", Args{1, 2}, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result=", result)
}
