package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		//var data [1024]byte
		arr := [1024]byte{}
		n, addr, err := listen.ReadFromUDP(arr[:])
		if err != nil {
			fmt.Println("read upd failed,err:", err)
		}
		fmt.Printf("data:%v addr%v count:%v\n", string(arr[:n]), addr, n)
		_, err = listen.WriteToUDP(arr[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed,err:", err)
		}
		continue
	}
}
