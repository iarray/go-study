package main

import "fmt"

var k int = 0

func main() {
	test2()
}

func test1(){

	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				//改为break+标签, 则不会只退出内层循环，而是直接退出外层循环了。
				//break LABEL1
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}


	fmt.Print("\nLABEL2 Start\n")

LABEL2:
	for i := 0; i <= 5; i++{
		k++
		for j := 0; j <= 5; j++ {
			if k > 2{
				break
			}
			if j == 4 {
				goto LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}

func test2(){
	// 	a := 1
	// 	goto TARGET // compile error
	// 	b := 9
	// TARGET:  
	// 	b += a
	// 	fmt.Printf("a is %v *** b is %v", a, b)
}