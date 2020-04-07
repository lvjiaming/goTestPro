package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "192.168.0.11:2001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("客户端")
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 注意：忽略错误
		log.Println("done")
		done <- struct{}{} // 向主Goroutine发出信号
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台goroutine完成
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
