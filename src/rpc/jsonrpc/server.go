package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MathMethod struct {
}

type Args struct {
	Val1, Val2 int
}

/* rpc要求至少有两个参数, 第一个参数是 */
func (*MathMethod) Add(arg *Args, reply *int) error {
	*reply = arg.Val1 + arg.Val2
	return nil
}

func main() {
	rpc.Register(new(MathMethod))
	addr, _ := net.ResolveTCPAddr("tcp4", ":8888")
	srv, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}

	for {
		client, err := srv.Accept()
		if err != nil {
			continue
		}

		//和tcp rpc唯一不同的地方
		go jsonrpc.ServeConn(client)
	}
}
