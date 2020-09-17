package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("start")

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sigterm:
			fmt.Println("recive signal")
			os.Exit(0)
		}
	}()

	for {
		fmt.Println("inloop")
		time.Sleep(time.Second * 2)
	}
}
