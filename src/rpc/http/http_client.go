package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	Val1, Val2 int
}

func main() {
	var result int
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	err = client.Call("MathMethod.Add", Args{1, 2}, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("MathMethod.Add(1,2)=", result)
}
