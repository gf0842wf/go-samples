package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"net/ttcp/types"
)

// network protocol
func SwitchNetProto(sess *types.Session) (resp []byte, err error) {
	if !sess.Shaked {
		// TODO: 服务器端发送握手消息给客户端
	}
	if sess.Encrypt {
		// TODO: 解密
	}
	return
}

// internal IPC
func SwitchIPCProto(sess *types.Session) (resp []byte, err error) {
	return
}

// 定时消息
func SwitchTMProto(sess *types.Session) (resp []byte, err error) {
	return
}
