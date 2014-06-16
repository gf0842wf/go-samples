package main

// TODO: 这里处理收到的消息

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

func HandleRequest(sess *types.Session, inChs chan []byte, outTag *Buffer) {
	// TODO: 在这里初始化session

	// the main message loop
	for {
		select {
		case msg, ok := <-inChs: // network protocol
			if !ok {
				return
			}
			fmt.Println("Data:", msg)
			// 对msg进行处理,并发送result数据给客户端...
			if sess.LoggedIn { // 已登录
				// TODO: 去game_proto  (判断是否是重进, 每把游戏有个唯一编号, 如果编号不存在了, 游戏就不存在了)
			} else { // 未登录
				// TODO: 去user_proto
			}
			result := msg
			err := outTag.Send(result)
			if err != nil {
				fmt.Println("Cannot send to client", err)
				return
			}
		// // outTag.ctrl <- true // 发送结束消息,注意这个消息发送后HandleRequest协程也要结束

		// // 其他消息, 如Session中的internal IPC, 控制消息, 定时器消息等
		// case msg := <-sess.MQ: // internal IPC
		// 	result := msg
		// 	err := outTag.Send(result)
		// 	if err != nil {
		// 		fmt.Println("Cannot send ipc response", err)
		// 		return
		// 	}
		// }
		case <-die:
			// TODD:
			return
		}
	}
}
