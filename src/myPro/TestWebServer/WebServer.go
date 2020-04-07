package TestWebServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
	count int
)


func TWSHandlerFunc()  {
	http.HandleFunc("/", handler) // 以“/”开头的 URL 链接在一起，代表所有的 URL 使用这个函数处理
	http.HandleFunc("/count", countHandler) // “/count”开头的连接请求
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%s", request.URL) // 将后面内容写入writer里
	})
	log.Fatal(http.ListenAndServe(":3030", nil))
}
/**
 w 是回复的消息，r是请求
 */
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()  // 锁（防止并发操作）
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}
	fmt.Fprintf(w, "Host = %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}

func countHandler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "count: %d", count)
	mu.Unlock()
}

// 加密的http服务，即https（ECHO）

func HttpsServer()  {
	//go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost(cert.pem, key.pem两个文件的生成)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprint(len("C语言中文网")))
		w.Write([]byte("C语言中文网"))
	})
	log.Fatal( http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
}

func EchoServer()  {

}