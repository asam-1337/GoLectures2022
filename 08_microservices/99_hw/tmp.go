package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	fmt.Println("vim-go")
	ctx := context.WithCancel(context.Background)

	select {
	case <-ctx.Done():
		fmt.Println("closed")
		return
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
	ctx.Done()
	time.Sleep(time.Second)

}
