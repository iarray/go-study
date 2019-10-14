package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := 0; i < 3; i++ {
		a := i
		go func() {
			for {
				time.Sleep(time.Second)
				fmt.Printf("GoRoutine%d Runing ...\n", a)
				select {
				case <-ctx.Done():
					fmt.Printf("GoRoutine%d Done !\n", a)
					return
				}
			}
		}()
	}

	time.Sleep(time.Duration(5) * time.Second)
	cancel()
	for {

	}
}
