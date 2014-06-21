package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"errors"
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/proto/login"
	"net/ttcp/types"
)

var uid_ interface{}
var uid uint32
var user *types.User

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
	k, _ := obj.K()
	switch k {
	case "SYS": // 握手登陆消息
		if handle, ok := login.SysProtoHandlers[kt]; ok {
			ack, err = handle(sess, &obj)
		}
	case "ROOM": // 房间消息
		uid_ = types.SessID2UID.Get(sess.ID)
		if uid_ == nil {
			err = errors.New("not logined")
			return
		}
		uid = uid_.(uint32)
		user = types.Users.Get(uid).(*types.User)
		fmt.Println(user)
	case "GAME": // 游戏中消息
		uid_ = types.SessID2UID.Get(sess.ID)
		if uid_ == nil {
			err = errors.New("not logined")
			return
		}
		uid = uid_.(uint32)
		user = types.Users.Get(uid).(*types.User)
		if !user.InGaming {
			err = errors.New("not gaming")
			return
		}
	default:
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
