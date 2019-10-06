package main

import "fmt"
 
func main(){
	i := 5
	p := &i

	*p = 6

	fmt.Println(p, i)

	var p2 *int = p

	*p2 = 8

	fmt.Println(p2, i)

	p3 := &p

	fmt.Println(p3, *p3, **p3)
	fmt.Printf("%T", p3)
	

	var p4 *int = nil
	*p4 = 0
	fmt.Println(p4, *p4)
}