package main

// TODO: 这里处理收到的消息

import (
	"fmt"
)

func HandleRequest(sess *Session, inChs chan []byte, outTag *Buffer) {
	// the main message loop
	for {
		select {
		case msg, ok := <-inChs: // network protocol
			if !ok {
				return
			}
			fmt.Println("Data:", msg)
			// 对msg进行处理...
			result := msg
			err := outTag.Send(result)
			if err != nil {
				fmt.Println("Cannot send to client", err)
				return
			}
			// 其他消息, 如Session中的internal IPC, 控制消息, 定时器消息等
		}
	}
}
