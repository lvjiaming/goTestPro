/**
 回音服务器（发送的数据原样返回）
利用windows的telnet客户端测试（需要打开：控制面板->程序->打开或关闭windows功能）
 */
package Telent

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
 连接服务器
 */
func server(address string, exitCh chan int)  {
	serverNet, err := net.Listen("tcp", address) // 监听函数（第一个参数为类型，第二个为地址）
	if err != nil {
		fmt.Print("err: ", err)
		exitCh <- 1 // 如果有错误，就退出
	}

	fmt.Print("已连接：", address)

	defer serverNet.Close() // 延迟退出

	for  {
		conn, err := serverNet.Accept() // 新连接未到来时，是阻塞的
		if err != nil {
			fmt.Println(err.Error())
			exitCh <- 2
			continue
		}
		go handler(conn, exitCh)
	}
}

func handler(conn net.Conn, exitCh chan int)  {
	fmt.Print("Session started: \n")
	reader := bufio.NewReader(conn)

	for  {
		str, err := reader.ReadString('\n') // 监听输入的字符串，以回车结束
		if err == nil {
			str := strings.TrimSpace(str) // 去掉字符串末尾的回车
			fmt.Print(str)
			if !processTelnetCommand(str, exitCh) {
				//fmt.Print("1")
				conn.Close()
				break
			}

			conn.Write([]byte(str + "\r\n")) // 原样返回
		} else {
			fmt.Print("session close")
			conn.Close()
			break
		}
	}
}

func processTelnetCommand(str string, exitCh chan int) bool {
	if strings.HasPrefix(str, "@close") { // 如果输入为@close则，关闭连接
		fmt.Print("connect closed\n")
		return false
	} else if strings.HasPrefix(str, "@shutdown") { //如果输入为@shutdown，关闭服务器
		fmt.Print("server closed\n")
		exitCh <- 4
		return false
	} else {
		fmt.Print(str)
		return true
	}
}

func TelnetSer(address string)  {
	exitch := make(chan int)

	go server(address, exitch)

	code := <- exitch
	fmt.Print("退出")
	os.Exit(code) // 打印code并退出
}
