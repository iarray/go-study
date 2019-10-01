package main

import "fmt"

func main()  {
	b := 'A' //int32, A, 65
	fmt.Printf("%T, %c, %d\n", b, b, b)

	var b2 byte = 'B' //uint8, B, 66
	fmt.Printf("%T, %c, %d\n", b2, b2, b2)

	c := '哈'  //int32, 哈, 54C8
	fmt.Printf("%T, %c, %X\n", c, c, c)

	var d int16 = '呵'  //int16, 呵, 5475
	fmt.Printf("%T, %c, %X\n", d, d, d)
	
	fmt.Printf("%c, %c \n", '\u54c8','\u5475')

	//默认字符类型是int32
	/*	
    判断是否为字母：unicode.IsLetter(ch)
    判断是否为数字：unicode.IsDigit(ch)
    判断是否为空白符号：unicode.IsSpace(ch)
	*/
}