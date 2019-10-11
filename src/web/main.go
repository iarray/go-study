package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
	"./session"
	_ "./session/provider/memory"
)

func main() {
	//simpleWeb()
	//customRouteHandler()
	//fileServer()
	useSession()
}

/* 简单输出 */
func simpleWeb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	//不设置handler就用默认的DefaultServeMux
	http.ListenAndServe(":8081", nil)
}

/* 自定义简易路由 */
type myRouteHandler struct {
}

func (*myRouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "hello world")
		return
	}
	if r.URL.Path == "/fuck" {
		fmt.Fprintf(w, "fuck you")
		return
	}
	if r.URL.Path == "/post" && r.Method == "POST" {
		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）

		//打印form表单
		fmt.Println(r.Form)
		//打印form某个值
		fmt.Println(r.Form["name"])
		//遍历表单
		for k, v := range r.Form {
			fmt.Fprintln(w, "key:", k)
			fmt.Fprintln(w, "val:", strings.Join(v, ","))
		}

		return
	}

	http.NotFound(w, r)
	return
}

func customRouteHandler() {
	handler := &myRouteHandler{}
	http.ListenAndServe(":8081", handler)
}

/* 文件上传 */
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method == "POST" {
		buf := make([]byte, 1<<10)
		hash := md5.New()
		//处理文件上传我们需要调用r.ParseMultipartForm，
		//里面的参数表示maxMemory，调用ParseMultipartForm之后，上传的文件存储在maxMemory大小的内存里面，
		//如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中
		//10mb, 2^10=1024
		r.ParseMultipartForm(10 << 20)
		fmt.Println("SetParseMultipartForm: 10mb")

		//我们可以通过r.FormFile获取上面的文件句柄
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		wfile, err := os.OpenFile("./upload/"+handler.Filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer wfile.Close()
		totalLen := handler.Size
		var totalRead int64
		for totalRead < totalLen {
			len, err := file.Read(buf)
			if err == nil {
				totalRead += int64(len)
				hash.Write(buf[:len])
				wfile.Write(buf[:len])
			}
		}

		fmt.Fprintln(w, "md5:", hex.EncodeToString(hash.Sum(nil)))
		fmt.Fprintln(w, "File:", wfile.Name())
	}

}

func fileServer() {
	os.Mkdir("./upload", 0777)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8081", nil)
}

/* session */
var globalSessionManager *session.SessionManager

func init() {
	var err error
	globalSessionManager, err = session.NewSessionManager("memory", "sessionId", 10*60)
	if err != nil{
		fmt.Println(err)
	}
}

func useSession(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		sess := globalSessionManager.SessionStart(w,r)
		fmt.Println(sess)
		val := sess.Get("Count")
		if val == nil{
			sess.Set("Count", 1)
			fmt.Fprintf(w, "你是第一次访问吧, 让我给你创建一个session, 稍后刷新看看~~")
		}else{
			sess.Set("Count", val.(int) + 1)
			fmt.Fprintf(w, "Count:%d\n", val)
			fmt.Fprintf(w, "SessionId:%s\n", sess.Id())
		}

	})
	http.ListenAndServe(":8081", nil)
}