package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime"
	"time"
)

import (
	// "zcodec"
	"net/ttcp/types"
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

	header := make([]byte, 4) // 假设包头只由length组成

	inChs := make(chan []byte, DEFAULT_INQUEUE_SIZE) // 存放解析好的一个完整消息的客户端发来的数据的channel
	ctrlCh := make(chan bool)                        // 控制停止发送消息的channel
	defer func() {
		close(ctrlCh)
		close(inChs)
	}()

	var sess types.Session

	outTag := NewBuffer(&sess, conn, ctrlCh) // 发送缓存

	go outTag.Start()                      // 开启发送缓存的协程
	go HandleRequest(&sess, inChs, outTag) // 开启处理收到消息的协程, 本程序主要的协程(其实outTag里有sess字段, 也有ctrl字段来控制停止发送)

	for {
		conn.SetReadDeadline(time.Now().Add(TCP_TIMEOUT * time.Second)) // 设置tcp读超时

		data, err := PreDecode(conn, header)
		if err != nil {
			break
		}

		// data初步解析好了, 得到了一个完整的消息, 送去继续处理
		select {
		case inChs <- data:
		case <-time.After(MAX_DELAY_IN * time.Second):
			fmt.Println("Pack timeout")
			return
		}
	}
	// TODO: 连接断开, 查询session的状态, 看是否在游戏中,如果不在,删除map中的session,
	// (如果在, 改变一下session状态为断开,)如果那局游戏结束,把掉线session在map中删除
}

func PreDecode(conn *net.TCPConn, header []byte) (data []byte, err error) {
	// header
	// --这个 ReadFull 非常好用, 作用是一直等到读取header大小的字节数为止
	n, err := io.ReadFull(conn, header)
	if err != nil {
		fmt.Println("Error recv header:", n, err)
	}

	// data
	// length := zcodec.ToUInt32(header, 62)
	length := binary.BigEndian.Uint32(header)
	data = make([]byte, length)
	n, err = io.ReadFull(conn, data)
	if err != nil {
		fmt.Println("Error recv msg:", n, err)
	}

	return
}
