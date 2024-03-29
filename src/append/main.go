package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//1. 删除位于索引 i 的元素 a := append(arr[:i], arr[i+1:]...)
	a := append(arr[:2], arr[2+1:]...)
	fmt.Println(a)   //[1 2 4 5 6 7 8 9 10]
	fmt.Println(arr) //[1 2 4 5 6 7 8 9 10 10]

	//2. 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = append(arr[:3], arr[6:]...) //4,5,6被移除
	fmt.Println(a)                  //[1 2 3 7 8 9 10]

	//3. 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
	fmt.Printf("&a=%p, typeof(a)=%T, len(a)=%d, cap(a)=%d\n", &a, a, len(a), cap(a))
	a = append(a, make([]int, 10)...)
	fmt.Printf("&a=%p, typeof(a)=%T, len(a)=%d, cap(a)=%d\n", &a, a, len(a), cap(a))

	//4. 在索引 i 的位置插入多个元素 ：a = append(a[:i], append([]T{x}, a[i:]...)...)
	//在索引2位置插入 11,12,13,   ...作用是把切片拆散传递
	a = append(a[:2], append([]int{11, 12, 13}, a[2:]...)...)
	fmt.Println(a)

	//5. 取出位于切片最末尾的元素
	x := a[len(a)-1]
	fmt.Println(x)
}
