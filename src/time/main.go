package main

<<<<<<< HEAD
import (
	"fmt"
	"time"
)

var week time.Duration

func main()  {
	t := time.Now()
	fmt.Println(t) // 2019-10-01 21:38:02.645699 +0800 CST m=+0.009996701
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 01.10.2019
	t = time.Now().UTC()
	fmt.Println(t) // 2019-10-01 13:38:02.7036848 +0000 UTC
	fmt.Println(time.Now()) // 2019-10-01 21:38:02.7036848 +0800 CST m=+0.067982501
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // 2019-10-08 13:38:02.7036848 +0000 UTC
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 01 Oct 19 13:38 UTC
	fmt.Println(t.Format(time.ANSIC)) // Tue Oct  1 13:38:02 2019
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 01 Oct 2019 13:38
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// 2019-10-01 13:38:02.7036848 +0000 UTC => 20191001
=======
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
>>>>>>> ab8c49282d59a228f86c8348879250262b6ca5f0
}