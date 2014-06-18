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

func HandleRequest(sess *types.Session, inChs chan []byte, outSender *types.SenderBuffer) {
	// TODO: 在这里初始化session
	sess.Sender = outSender
	types.Sessions.Set(types.SessID, sess)
	// stats.Sessions.Set(stats.SessID, sess)
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
				// TODO: 去game_proto
			} else { // 未登录
				// TODO: 去user_proto
				// 登陆时要判断是否是重进, 每把游戏有个唯一编号, 如果编号不存在了, 游戏就不存在了
				// 如果是重进,则要把之前的sess copy到新sess中,并在map中删除,这个添加到map..并且 outSender 也要这么处理
				// 要关闭旧连接(如果旧连接还在的话,此时就是异地登陆,要踢掉自己)
			}
			result := msg
			err := outSender.Send(result)
			if err != nil {
				fmt.Println("Cannot send to client", err)
				return
			}
		case msg := <-sess.MQ: // internal IPC
			// TODO: 去ipc_proto
			fmt.Println(msg)
			result := []byte("test")
			err := outSender.Send(result)
			if err != nil {
				fmt.Println("Cannot send ipc response", err)
				return
			}
		// 其他消息, 如Session中的internal IPC, 控制消息, 定时器消息等
		case <-die:
			// TODD: 清理信息
			return
		}
	}
}
