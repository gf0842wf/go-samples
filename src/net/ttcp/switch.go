package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"errors"
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/proto/auth"
	"net/ttcp/proto/room"
	"net/ttcp/types"
)

// network protocol
func SwitchNetProto(user *types.User, data []byte) (ack []byte, err error) {
	var obj proto.Msg
	err = user.Sess.Coder.Decode(data, &obj)
	if err != nil {
		fmt.Println("SwitchNetProto decode err:", err.Error())
		return
	}
	kt, err := obj.KT()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	k, _ := obj.K()
	switch k {
	case "AUTH": // 握手登陆消息
		if handle, ok := auth.AuthProtoHandlers[kt]; ok {
			ack, err = handle(user, &obj)
		}
	case "ROOM": // 房间消息
		if !user.Logined {
			err = errors.New("not logined")
			return
		}
		if handle, ok := room.RoomProtoHandlers[kt]; ok {
			ack, err = handle(user, &obj)
		}
	case "GAME": // 游戏中消息
		if !user.InGaming {
			err = errors.New("not in gaming")
			return
		}
		// TODO: 每个user有deskno号,然后可以直接分发到对于的desk goroutine的MQ
		// 用一个单独的协程来分发消息吧
	default:
	}

	return
}

// internal IPC
func SwitchIPCProto(user *types.User, data []byte) (ack []byte, err error) {
	return
}

// 定时消息
func SwitchTMProto(user *types.User, data []byte) (ack []byte, err error) {
	return
}
