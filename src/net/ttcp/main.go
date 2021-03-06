package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

import (
	"net/ttcp/codec"
	"net/ttcp/match"
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

	user := types.NewUser()
	sess := types.NewSession()
	user.Sess = sess
	coder := codec.NewCoder()
	sess.Coder = coder

	outSender := types.NewSenderBuffer(sess, conn, ctrlCh) // 发送缓存

	go outSender.Start()                     // 开启发送缓存的协程
	go HandleRequest(user, inChs, outSender) // 开启处理收到消息的协程, (也有ctrl字段来控制停止发送)

	for {
		conn.SetReadDeadline(time.Now().Add(TCP_TIMEOUT * time.Second)) // 设置tcp读超时
		data, err := codec.PreDecode(conn, header)
		if err != nil {
			fmt.Println(err.Error())
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
	// TODO: 如果那局游戏结束,把掉线user在map中删除
	Clear(user)

}

func Clear(user *types.User) {
	user.IsActive = false
	user.Sess.Coder.Shaked = false
	if !user.InGaming {
		match.UserLeave(user, 0)
		types.Users.Delete(user.ID)
	}
	fmt.Println("Clear user:", user.ID)
}
