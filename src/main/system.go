package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func a(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done!")
			return
		default:
			fmt.Println("inloop")
			time.Sleep(time.Second * 2)
		}
	}
}
func main() {
	fmt.Println("start")

	ctx, cancel := context.WithCancel(context.Background())

	go a(ctx)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		fmt.Println("recive signal")
		cancel()
		fmt.Println("recived signal")
		os.Exit(0)
	}
}
