package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	tcpClientConnectWithTimeout()
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		tcpClient()
	}()

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		tcpClientDelayWrite()
	}()

	tcpServer()
}

func tcpClient() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	conn, err := net.DialTCP("tcp", nil, addr)
	handlerErr(err)
	len, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	handlerErr(err)
	fmt.Println("Write Success len = ", len)

	var buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	handlerErr(err)
	fmt.Println("Read Success")
	fmt.Println(string(buf))
	conn.Close()
}

//测试超时发送
func tcpClientDelayWrite() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	conn, err := net.DialTCP("tcp", nil, addr)
	handlerErr(err)

	time.Sleep(time.Duration(2) * time.Second)
	len, err := conn.Write([]byte("PING"))
	handlerErr(err)
	fmt.Println("Write Success len = ", len)

	conn.Close()
}

//带超时的连接
func tcpClientConnectWithTimeout() {
	conn, err := net.DialTimeout("tcp4", "127.0.0.1:8082", time.Duration(500)*time.Millisecond)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
}

func tcpServer() {
	addr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:8082")
	listener, err := net.ListenTCP("tcp", addr)
	handlerErr(err)
	for {
		client, _ := listener.Accept()
		go onClientConnect(client)
	}
}

func onClientConnect(client net.Conn) {
	//设置建立连接的超时时间，客户端和服务器端都适用，当超过设置时间时，连接自动关闭。
	//func (c *TCPConn) SetReadDeadline(t time.Time) error
	//func (c *TCPConn) SetWriteDeadline(t time.Time) error

	//读取等待超过1秒终止
	client.SetReadDeadline(time.Now().Add(time.Duration(1) * time.Second))

	buf := make([]byte, 1024)
	_, err := client.Read(buf)
	handlerErr(err)
	fmt.Println(string(buf))
	client.Write([]byte("HTTP/1.0 200 OK\r\n\r\n"))
	//或者这样
	fmt.Fprint(client, "Date: Fri, 11 Oct 2019 08:14:59 GMT")
}

func handlerErr(err error) {
	if err != nil {
		panic(err)
	}
}
