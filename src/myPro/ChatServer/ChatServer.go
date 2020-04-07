/**
聊天服务器
 */
package ChatServer

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- interface{} // 单向通道(只能发送)
var (
	message = make(chan interface{})
	enter = make(chan client)
	leave = make(chan client)
)

func ChatServer()  {
	listener, err := net.Listen("tcp", "192.168.0.11:2001")
	if err != nil {
		fmt.Println(err)
	}
	go registerEvent()
	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handler(client)
	}
}

func registerEvent()  {
	clients := make(map[client]bool)
	for  { // 一定要在for循环里循环去接受chan
		select {
		case msg := <-message:
			//fmt.Print(clients)
			for cli, _ := range clients {
				cli <- msg
			}
		case client := <-enter:
			clients[client] = true
		case client := <-leave:
			//clients[client] = false
			delete(clients, client)
			close(client)
		}
	}
}

func handler(conn net.Conn)  {
	fmt.Print("有客户端连接\n")
	ch := make(chan interface{})
	go repMsg(conn, ch)
	info := conn.RemoteAddr().String()
	ch <- "欢迎" + info
	message <- info + "上线"
	enter <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Print(input.Text())
		message <- info + ": " + input.Text()
	}
	leave <- ch
	message <- info + "下线"
	fmt.Print("有客户端断开连接\n")
	conn.Close()
}

func repMsg(conn net.Conn, ch chan interface{})  {
	for msg := range ch{
		fmt.Fprintln(conn, msg)
	}
}