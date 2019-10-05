package main

import(
	"time"
	"fmt"
)

func main()  {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		
	}
	end := time.Now()
	span := end.Sub(start)
	fmt.Printf("cost=%s\n, typeof(span)=%T", span, span)
}