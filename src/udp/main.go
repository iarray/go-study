package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			udpClient()
			udpBoardCast()
		}
	}()

	udpServer()
}

func udpClient() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8082")
	conn, err := net.DialUDP("udp", nil, addr)
	checkErr(err)

	len, _ := conn.Write([]byte("ping"))
	fmt.Println("Write Len=", len)

	len, _ = conn.Write([]byte("hello"))
	fmt.Println("Write Len=", len)
}

//发送广播
func udpBoardCast() {
	boardcast, _ := net.ResolveUDPAddr("udp", "255.255.255.255:8082")
	conn, err := net.DialUDP("udp", nil, boardcast)
	checkErr(err)

	len, _ := conn.Write([]byte("boardcast ping"))
	fmt.Println("Write Len=", len)

	len, _ = conn.Write([]byte("boardcast hello"))
	fmt.Println("Write Len=", len)

	//写失败
	conn.WriteToUDP()
}

func udpServer() {
	addr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:8082")
	listener, err := net.ListenUDP("udp", addr)
	checkErr(err)

	for {
		buf := make([]byte, 1024)
		len, udpClientAddr, err2 := listener.ReadFromUDP(buf)
		if err2 != nil {
			fmt.Println("Read Fail", err2)
			continue
		}
		go onRecieveData(udpClientAddr, buf, len)
	}
}

func onRecieveData(udpClientAddr *net.UDPAddr, buf []byte, len int) {
	fmt.Printf("onRecieveData From %s\n", udpClientAddr.String())
	fmt.Println(string(buf))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
