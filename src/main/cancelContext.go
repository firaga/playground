package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
			//go reqTask(ctx, "worker3")
		}
	}
}

func main() {
	defer func() {
		fmt.Println("defer of main")
	}()
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")
	//time.Sleep(3 * time.Second)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		fmt.Println("recive signal")
		cancel()
		//time.Sleep(3 * time.Second)
		fmt.Println("recived signal")
		//os.Exit(0)
	}
	os.Exit(0)
	//time.Sleep(3 * time.Second)
}
