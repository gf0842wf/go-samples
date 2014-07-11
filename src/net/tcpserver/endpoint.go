package tcpserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

/* tcp Endpoint:
创建 EndPoint 对象后需要调用 Init 和 Start 方法
对外暴露的接口:
1.使用 PutData 发送
2.读取 RecvBox 管道处理收到的消息
3.设置 OnConnectionLost 回调函数来处理断开连接, 可以为nil
*/

type EndPoint struct {
	Conn *net.TCPConn

	SendBox chan []byte // 发送缓冲管道
	RecvBox chan []byte // 接收缓冲管道
	Ctrl    chan bool   // 控制结束 EndPoint 所有协程的

	Heartbeat int64 // 心跳超时(s), < 0表示不设置心跳

	OnConnectionLost interface{} // func(err error)  回调, err: 断开错误信息
}

func (ep *EndPoint) Init(conn *net.TCPConn, heartbeat int64, sendBufSize int, recvBufSize int, OnConnectionLost interface{}) {
	ep.Conn = conn
	ep.Heartbeat = heartbeat
	ep.Ctrl = make(chan bool)
	ep.SendBox = make(chan []byte, sendBufSize)
	ep.RecvBox = make(chan []byte, recvBufSize)
	if OnConnectionLost == nil {
		ep.OnConnectionLost = ep.onConnectionLost
	} else {
		ep.OnConnectionLost = OnConnectionLost
	}
}

func (ep *EndPoint) PutData(data []byte) {
	ep.SendBox <- data
}

func (ep *EndPoint) recvData() {
	header := make([]byte, 4)

	var err error

	for {
		if ep.Heartbeat > 0 {
			ep.Conn.SetReadDeadline(time.Now().Add(time.Duration(ep.Heartbeat) * time.Second))
		}
		_, err = ep.RawRecv(header)
		if err != nil {
			break
		}
	}
	ep.OnConnectionLost.(func(err error))(err)
}

// 如果封包方式不同,需要修改这个函数
func (ep *EndPoint) RawRecv(header []byte) (n int, err error) {
	// header
	// --这个 ReadFull 非常好用, 作用是一直等到读取header大小的字节数为止
	n, err = io.ReadFull(ep.Conn, header)
	if err != nil {
		err = errors.New("[EP] Error recv header:" + strconv.Itoa(n) + ":" + err.Error())
		return
	}

	// data
	length := binary.BigEndian.Uint32(header)
	data := make([]byte, length)
	n, err = io.ReadFull(ep.Conn, data)
	if err != nil {
		err = errors.New("[EP] Error recv msg:" + strconv.Itoa(n) + ":" + err.Error())
		return
	}
	ep.onData(data)

	return
}

func (ep *EndPoint) onData(data []byte) {
	ep.RecvBox <- data
}

func (ep *EndPoint) onConnectionLost(err error) {
	fmt.Println("[EP] Connection Lost:", err.Error())
	ep.Ctrl <- false
}

func (ep *EndPoint) sendData() {
	header := make([]byte, 4)
	for {
		select {
		case data := <-ep.SendBox:
			ep.RawSend(header, data)
		case <-ep.Ctrl:
			defer close(ep.SendBox)
			defer close(ep.RecvBox)
			// 准备关闭连接, 要发完剩下的消息
			for data := range ep.SendBox {
				ep.RawSend(header, data)
			}
			ep.Conn.Close()

			fmt.Println("[EP] Close connection:", ep.Conn.LocalAddr)

			return
		}
	}
}

// 如果封包方式不同,需要修改这个函数
func (ep *EndPoint) RawSend(header []byte, msg []byte) {
	length := len(msg)
	binary.BigEndian.PutUint32(header, uint32(length))
	data := append(header, msg...)
	n, err := ep.Conn.Write(data)
	if err != nil {
		fmt.Println("[EP] Error send reply, bytes:", n, "reason:", err)
		return
	}
}

func (ep *EndPoint) Start() {
	go ep.recvData()
	go ep.sendData()
}
