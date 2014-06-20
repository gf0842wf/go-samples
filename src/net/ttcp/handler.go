package main

// 这里处理收到的消息

import (
	"fmt"
)

import (
	"net/ttcp/types"
)

// 结束 HandleRequest 协程的channel
var die chan bool // for server close
func init() {
	die = make(chan bool)
}

func HandleRequest(sess *types.Session, inChs chan []byte, outSender *types.SenderBuffer) {
	// TODO: 在这里更新session的时间等信息
	sess.Sender = outSender
	// the main message loop
	for {
		select {
		case msg, ok := <-inChs: // network protocol
			if !ok {
				return
			}
			fmt.Println("Data:", string(msg))
			result, err := SwitchNetProto(sess, msg)
			if err != nil {
				// 断开连接
				fmt.Println(err.Error())
				sess.Disconnect()
				return
			}
			if result != nil {
				err = outSender.Send(result)
				if err != nil {
					fmt.Println("Cannot send to client:", err)
					return
				}
			}
		case msg := <-sess.MQ: // internal IPC
			// TODO: 去ipc_proto
			fmt.Println("MQ:", msg)
			result := []byte("test")
			err := outSender.Send(result)
			if err != nil {
				fmt.Println("Cannot send ipc response:", err)
				return
			}
		// 其他消息, 如Session中的internal IPC, 控制消息, 定时器消息等
		case <-die:
			// TODD: 清理信息
			return
		}
	}
}
