package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"runtime"
)

const (
	ADDR = ":8888"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Caught panic in main()")
			panic(x)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU()) // 开启多核

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ADDR)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server start:", ADDR)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Caught panic in handleClient")
			panic(x)
		}
	}()

	var header []byte
	var method, url string // "GET", "http://www.baidu.com/"
	var host string        // "www.baidu.com:80"

	data := make([]byte, 128)
	buf := make([]byte, 0)
	for {
		conn.SetReadDeadline(time.Now().Add(180 * time.Second)) // 设置tcp读超时
		n, err := conn.Read(data)
		buf = append(buf, data[:n]...)
		if err != nil {
			fmt.Println("Read err")
			break
		}
		pos := bytes.IndexAny(buf, "\r\n\r\n")
		if pos == -1 {
			continue
		}
		// 头部读完
		header = buf[:index]
		buf = buf[index+4:]
		// 解析头部

	}

}
