package main

import(
	"bytes"
	"strconv"
	"fmt"
)

func main(){
	var buf bytes.Buffer
	for i := 0; i < 10; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(i))
	}
	fmt.Printf("type=%T, value=%s\n", buf, buf.String())
}