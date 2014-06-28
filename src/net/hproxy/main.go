package main

import (
	"bytes"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	ADDR = ":8081"
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

	gotHeader := false
	var header []byte
	var ctx []byte
	var method, url []byte // "GET", "http://www.baidu.com/"
	var host []byte        // "www.baidu.com:80"

	var posCtxLen int = -1
	var ctxLen int = 0

	data := make([]byte, 128)
	buf := make([]byte, 0)
	for {
		conn.SetReadDeadline(time.Now().Add(180 * time.Second)) // 设置tcp读超时
		n, err := conn.Read(data)
		buf = append(buf, data[:n]...)
		if err != nil {
			fmt.Println("Svr read err")
			break
		}
		if !gotHeader {
			pos := bytes.Index(buf, []byte("\r\n\r\n"))
			if pos == -1 {
				continue
			}
			// 头部读完
			header = buf[:pos]
			buf = buf[pos+4:]
			// 解析头部
			firstPos := bytes.Index(header, []byte("\r\n"))
			firstLine := header[:firstPos]
			mu := bytes.Split(firstLine, []byte(" "))
			if len(mu) == 2 || len(mu) == 3 {
				method = mu[0]
				url = mu[1]
				fmt.Println("Mu:", string(method), string(url))
			} else {
				fmt.Println("Mu error", mu)
			}
			host = bytes.TrimRight(bytes.TrimLeft(url, "http://"), "/")
			posCtxLen = bytes.Index(buf, []byte("Content-Length: "))
			gotHeader = true
		}
		if posCtxLen == -1 {
			// 没有内容
			header = append(header, []byte("\r\n\r\n")...)
			go forward(conn, header, host)
		} else { // 有内容
			ctxLine := header[posCtxLen:50]
			posCtx := bytes.Index(ctxLine, []byte("\r\n"))
			ctxLen, _ = strconv.Atoi(string(ctxLine[len("Content-Length: "):posCtx]))
			if len(buf) < ctxLen {
				continue
			} else {
				ctx = buf[:ctxLen]
				buf = buf[ctxLen:]
				header = append(header, []byte("\r\n\r\n")...)
				msg := append(header, ctx...)
				fmt.Println(string(host), msg[0:1], posCtxLen)
				go forward(conn, msg, host)
			}
		}
	}
	fmt.Println("Svr connection lost")
}

func forward(svrConn *net.TCPConn, msg []byte, host []byte) {
	defer svrConn.Close()
	shost := string(host)
	if strings.Count(shost, ":") == 0 {
		shost += ":80"
	}
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", shost)
	// if err != nil {
	// 	fmt.Println("shost:", shost)
	// 	fmt.Println("Cli addr parse error")
	// 	return
	// }
	cliConn, err := net.Dial("tcp", shost)
	if err != nil {
		fmt.Println("Cli dial:", err.Error())
		return
	}
	defer cliConn.Close()
	n, e := cliConn.Write(msg)
	if e != nil {
		fmt.Println("Cli write err:", n, e.Error())
		return
	}
	data := make([]byte, 128)
	for {
		n, e := cliConn.Read(data)
		if e != nil {
			fmt.Println("Cli read err:", e.Error())
			break
		}
		sn, se := svrConn.Write(data[:n])
		if se != nil {
			fmt.Println("Svr write err:", sn, se.Error())
			break
		}
	}
	fmt.Print("Cli connection lost")
}
