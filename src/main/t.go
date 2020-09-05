package main

import (
	"fmt"
	"strings"
)

func main() {
	s :="\r\nhello \r world\nhaha\r\n"
	a:=strings.Trim(s,"\n")
	fmt.Println(len(s))
	fmt.Println(len(a))
}
