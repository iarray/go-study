package main

import "fmt"

func main(){
   s:= "hello 小米"	
   //因为go的顶层string是[]byte, 对于unicode字符占两字节, 而go默认编码是utf-8, 一般utf8英文占1字节,中文占3字节
   //所以长度是 hello(5) + 空格(1) + 小米(6) = 
   fmt.Printf("len(s)=%d\n", len(s)) //len=12 fuck!
   fmt.Printf("len(小)=%d\n", len("小")) 
   fmt.Printf("len(米)=%d\n", len("米"))

   //所以计算字符串的长度应该先将字符串转为[]rune 再计算
   fmt.Printf("len([]rune(s))=%d\n", len([]rune(s))) //len=8 yes!
   fmt.Printf("len([]int32(s))=%d\n", len([]int32(s))) //len=8 yes!
   
}