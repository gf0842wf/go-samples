package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/sys_proto"
	"net/ttcp/types"
)

// network protocol
func SwitchNetProto(sess *types.Session, data []byte) (ack []byte, err error) {
	var obj proto.Msg
	err = sess.Coder.Decode(data, &obj)
	if err != nil {
		fmt.Println("SwitchNetProto decode err:", err.Error())
		return
	}
	kt, err := obj.KT()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch kt {
	case "SYS.PRESHAKE", "SYS.ACKSHAKE": // 等等
		if handle, ok := sys_proto.SysProtoHandlers[kt]; ok {
			ack, err = handle(sess, &obj)
		}
	default:
		//  这个是去游戏消息的

	}
	if !sess.Coder.Shaked {
		// 这个时机断开连接: 因为已经第一次收数据了
	}

	return
}

// internal IPC
func SwitchIPCProto(sess *types.Session, data []byte) (ack []byte, err error) {
	return
}

// 定时消息
func SwitchTMProto(sess *types.Session, data []byte) (ack []byte, err error) {
	return
}
