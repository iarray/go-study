package main

import (
	"net"
	"fmt" 
	"bytes"
)

func main()  {
	l,err := net.Listen("tcp",":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	for{

		client, err := l.Accept()
		if err != nil {
			continue
		}

		go handlerClient(client)
	}

}

func handlerClient(client net.Conn)  {
	var buf = make([]byte, 1024, 1024)
	len , err := client.Read(buf)
	if err != nil{
		return
	}

	fmt.Printf("read len=%d\n", len)
	str := string(buf[:len])
	fmt.Println(str)

	var method, host, proto string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host, &proto)
	 
	fmt.Printf("method =%s, host=%s, proto=%s \n", method, host, proto)


	client.Write([]byte("Hello"))
}