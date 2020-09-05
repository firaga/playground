package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		//fmt.Println("input: ", len(input))
		inputInfo := strings.Trim(input, "\r\n")
		//fmt.Println("inputInfo:", len(inputInfo))
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		byteInputInfo := []byte(inputInfo)
		fmt.Println(byteInputInfo)
		fmt.Println(cap(byteInputInfo))
		fmt.Println(len(byteInputInfo))
		_, err := conn.Write(byteInputInfo)
		if err != nil {
			return
		}

		buf := [3]byte{}
		fmt.Println(reflect.TypeOf(buf))
		fmt.Println(cap(buf))
		fmt.Println(len(buf))
		fmt.Println(reflect.TypeOf(buf[:]))
		s_buf := buf[:]
		fmt.Println(cap(s_buf))
		fmt.Println(len(s_buf))
		n, err := conn.Read(s_buf)
		fmt.Println(n)
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println("recved", string(buf[:n]))
	}
}
