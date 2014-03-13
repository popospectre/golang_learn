package main

import (
	"fmt"
	"net"
	//"os"
)

func main() {

	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage ID: ", os.Args[0], "host")
	// 	os.Exit(1)
	// }

	service := "127.0.0.1:8088"

	conn, err := net.Dial("tcp", service)
	if err != nil {
		fmt.Println("client connect err:", err.Error())
		return
	}

	defer conn.Close()

	var msg [1024]byte
	copy(msg[:], "hello world")
	_, err = conn.Write(msg[0:1023])

	if err != nil {
		fmt.Println("client write err:", err.Error())
		return
	}
	fmt.Println("client write data: ", string(msg[:]))

	var buf [1024]byte
	_, err = conn.Read(buf[0:])

	fmt.Println("client recv data: ", string(buf[:]))

}
