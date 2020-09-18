package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 结论
// 主协程结束,所有协程强制结束
// 儿子协程结束,孙子协程继续执行
func main() {
	defer func() {
		fmt.Println("exit main")
	}()
	fmt.Println("main start")
	var start chan struct{}
	start = make(chan struct{})
	go func() {
		defer func() {
			fmt.Println("go 1 exit")
		}()
		start <- struct{}{}
		go func() {
			name := "subsub.txt"
			file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			fmt.Println("go 2")
			i := 0
			file.WriteString(strings.Repeat("-", 40))
			file.WriteString("\n")
			for ; i < 6; i++ {
				file.WriteString(fmt.Sprintf("hello %d \n", i))
				fmt.Println("go 2 loop", i)
				time.Sleep(time.Second)
			}
		}()
		name := "sub.txt"
		file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		fmt.Println("go 1")
		i := 0
		file.WriteString(strings.Repeat("-", 40))
		file.WriteString("\n")
		for ; i < 6; i++ {
			file.WriteString(fmt.Sprintf("hello %d \n", i))
			fmt.Println("go 1 loop", i)
			//time.Sleep(time.Second)
		}
	}()
	<-start
	time.Sleep(time.Second * 10)
	fmt.Println("main end")
}
