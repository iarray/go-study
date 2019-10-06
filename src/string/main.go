package main

<<<<<<< HEAD
import (
	"fmt" 
	"strconv"
	"strings"
)

func main()  {
	
	//是否有前缀
	str := "Thank you"
	var b bool
	b = strings.HasPrefix(str, "Th")
	fmt.Println(b);

	//是否有后缀
	b = strings.HasSuffix(str, "you")
	fmt.Println(b);

	//是否包含
	b = strings.Contains(str, "ank")
	fmt.Println(b);

	//字符串索引查找
	//strings.Index(s, str string) int
	//strings.LastIndex(s, str string) int
	//如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：
	//strings.IndexRune(s string, r rune) int

	//替换
	//strings.Replace(str, old, new, n) string

	//统计出现次数, Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：
	//strings.Count(s, str string) int

	//生成重复字符, Repeat 用于重复 count 次字符串 s 并返回一个新的字符串：
	//strings.Repeat(s, count int) string

	//大小写转换
	//strings.ToLower(s) string
	//strings.ToUpper(s) string

	//删除首尾空格
	//strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号

	//分割字符串
	//strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

	//通过分割符号拼接
	//strings.Join(sl []string, sep string) string

	//从字符串读取自己
	r := strings.NewReader("abcdefghijklmn")
	var buf []byte
	buf = make([]byte, 5)
	r.Read(buf)
	fmt.Println(buf);


	//类型转换
	
	//字符转整形
	i, err := strconv.Atoi("123456")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(i);

	//返回数字 i 所表示的字符串类型的十进制数。
	//strconv.Itoa(i int) string

	//将字符串转换为 float64 型。
	//strconv.ParseFloat(s string, bitSize int) (f float64, err error) 

	/*
	strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，
	其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），
	prec 表示精度，
	bitSize 则使用 32 表示 float32，用 64 表示 float64。
	*/
=======
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
   
>>>>>>> ab8c49282d59a228f86c8348879250262b6ca5f0
}