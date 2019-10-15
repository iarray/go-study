package main

import (
	"net/http"
	"net/rpc"
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
	//rpc.HandleHTTP函数把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}

}
