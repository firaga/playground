package main

import (
	"bufio"
	"fmt"
	"net"
)

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		fmt.Println("after reader read")
		if err != nil {
			fmt.Println("read from client failed,err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
	}

	for {
		fmt.Println("forloop")
		conn, err := listen.Accept()
		fmt.Println("accept")
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		go Process(conn)
	}
}
