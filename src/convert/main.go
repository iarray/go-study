package main

import (
	"fmt" 
	"strconv"
)

func main()  {
	i, err := strconv.Atoi("123456")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(i);
}