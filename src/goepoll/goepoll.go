package goepoll

import (
	"fmt"
	"net"
	"time"
)

const (
	MAX_CONN_NUM = 5
)

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			//println("Error reading:", err.Error())
			return
		}

		fmt.Println("server recv :", string(buf[:]))
		//send reply
		_, err = conn.Write(buf)
		if err != nil {
			//println("Error send reply:", err.Error())
			return
		}
	}
}

// 启动epoll
// addr: ip和port, 如127.0.0.1:8088
// err: 返回错误
func RunEpoll(addr string) (err error) {
	listener, err1 := net.Listen("tcp", addr)
	if err1 != nil {
		fmt.Println("error listening:", err1.Error())
		err = err1
		return err
	}

	defer listener.Close()
	fmt.Println("goepool runnin...")

	var cur_conn_num int = 0 // 当前连接数
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		for _ = range time.Tick(1e8) {
			fmt.Println("goepoll cur conn num:", cur_conn_num)
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				ch_conn_change <- 1
				EchoFunc(conn)
				ch_conn_change <- -1
			}
		}()
	}

	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("Error accept:", err1.Error)
			err = err1
			return err
		}
		conn_chan <- conn
	}
}
