package main

import(
	"net"
	"fmt"
	"os"
)

func main()  {
	ip := net.ParseIP("192.168.101.111")
	if ip != nil {
		fmt.Println(ip.String())
	}
	os.Exit(0)
}