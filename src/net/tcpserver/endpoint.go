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
*/

type EndPoint struct {
	Conn      *net.TCPConn
	SendBox   chan []byte // 发送缓冲管道
	Ctrl      chan bool   // 控制结束 EndPoint 所有协程的
	Heartbeat int64       // 心跳超时(s), < 0表示不设置心跳

	OnData           func(data []byte) // 回调, data: 解析后的消息
	OnConnectionLost func(err error)   // 回调, err: 断开错误信息
}

func (ep *EndPoint) Init(conn *net.TCPConn, heartbeat int64, bufSize int, OnData func(data []byte), OnConnectionLost func(err error)) {
	ep.Conn = conn
	ep.Heartbeat = heartbeat
	ep.Ctrl = make(chan bool)
	ep.SendBox = make(chan []byte, bufSize)
	ep.OnData = OnData
	ep.OnConnectionLost = OnConnectionLost
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
	ep.OnConnectionLost(err)
}

// 如果封包方式不同,需要修改这个函数
func (ep *EndPoint) RawRecv(header []byte) (n int, err error) {
	// header
	// --这个 ReadFull 非常好用, 作用是一直等到读取header大小的字节数为止
	n, err = io.ReadFull(ep.Conn, header)
	if err != nil {
		err = errors.New("[EndPoint] Error recv header:" + strconv.Itoa(n) + ":" + err.Error())
		return
	}

	// data
	length := binary.BigEndian.Uint32(header)
	data := make([]byte, length)
	n, err = io.ReadFull(ep.Conn, data)
	if err != nil {
		err = errors.New("[EndPoint] Error recv msg:" + strconv.Itoa(n) + ":" + err.Error())
		return
	}
	ep.OnData(data) // go OnData(data)

	return
}

func (ep *EndPoint) sendData() {
	header := make([]byte, 4)
	for {
		select {
		case data := <-ep.SendBox:
			ep.RawSend(header, data)
		case <-ep.Ctrl:
			defer close(ep.SendBox)
			// 准备关闭连接, 要发完剩下的消息
			for data := range ep.SendBox {
				ep.RawSend(header, data)
			}
			ep.Conn.Close()

			fmt.Println("[EndPoint] Close connection:", ep.Conn.LocalAddr)

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
		fmt.Println("[EndPoint] Error send reply, bytes:", n, "reason:", err)
		return
	}
}

func (ep *EndPoint) Start() {
	go ep.recvData()
	go ep.sendData()
}
