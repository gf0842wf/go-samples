package main

// 转接handler里network protocol 和 internal IPC 等

import (
	"errors"
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/sys_proto"
	"net/ttcp/types"
)

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
	switch kt {
	case "SYS.PRESHAKE", "SYS.ACKSHAKE":
		if handle, ok := sys_proto.SysProtoHandlers[kt]; ok {
			ack, err = handle(sess, &obj)
		}
	case "SYS.LOGIN":
		if !sess.Coder.Shaked {
			err = errors.New("not shaked")
		}
		if handle, ok := sys_proto.SysProtoHandlers[kt]; ok {
			ack, err = handle(sess, &obj)
		}
	default:
		if !sess.Coder.Shaked {
			err = errors.New("not shaked")
			return
		}
		uid_ := types.SessID2UID.Get(sess.ID)
		if uid_ == nil {
			err = errors.New("not logined")
			return
		}
		uid = uid_.(uint32)
		user = types.Users.Get(uid).(*types.User)
		fmt.Println(user)
		//  这个是去游戏消息的

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
