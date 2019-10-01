package main

import (
	log "fmt"
)

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9}
	s1 := arr[:2]	 //len=2, cap=2
	s2 := arr[1:2:3] //len=1, cap=3

	s1[0] = 12
	s1[1] = 11
	//s1[2] = 0 //err
	//2,1
	log.Println(s1);

	s2[0] = 14
	//s2[1] = 5  //err
	//s2[2] = 6  //err
	log.Println(s2)

	s3 := append(s2,10) //len=2, cap=3
	//s3[0] = 24
	//s3[1] = 25
	//s3[2] = 6  //err
	log.Println(s2)
	log.Println(s3)


	s3 = append(s3,11) //len=3, cap=3
	s3[2] = 26 
	log.Println(s3)

	log.Println(arr)
}