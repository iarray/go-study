package hello

var(
	PI float32
)

/*
init不能显示调用， 每个包完成初始化后自动执行init，并且执行优先级比 main 函数高。
*/
func init()  {
	PI = 3.14
}

func GetName()string{
	return "aaa"
}