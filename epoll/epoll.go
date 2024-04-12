package epoll

import (
	"fmt"
	"net"
	"time"
)

func RunTCP() {
	// TCP监听器

	go func() {
		tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:30081") // bind 阶段， 绑定ip和端口
		if err != nil {
			panic(err)
		}
		l, err := net.ListenTCP("tcp", tcpaddr) // 调用 epoll_create // listen阶段， 监听IO TCP请求
		defer func() {
			err := l.Close() // close阶段
			panic(err)
		}()
		if err != nil {
			panic(err)
		}
		// 自旋等待
		for {
			// 什么时候会触发:
			// 1. 全局监控任务
			// 2. GMP调度
			// 3. GC start-the-world
			conn, _ := l.Accept() // 调用epoll_wait阻塞(gopark), 避免自选浪费CPU; 调用epoll_ctl去处理真正关心的数据 // 等待客户端连接, 如果没有就阻塞, 避免空占CPU
			go serve(conn)
		}
	}()

	time.Sleep(2 * time.Second)
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:30081") // 客户端的Connect阶段, 或者是Dial阶段
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("hello world"))
	}()

	for {
		time.Sleep(1 * time.Second)
	}

}

func serve(conn net.Conn) {
	fmt.Println("in server")
	defer conn.Close()
	buf := make([]byte, 1024)
	num, err := conn.Read(buf) // 读不到的时候也会阻塞(gopark)

	if err == nil {
		fmt.Println("num: ", num)
		fmt.Println("buf: ", string(buf))
	}

}
