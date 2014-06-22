package auth

// 心跳

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_nop(user *types.User, obj *proto.Msg) (resp []byte, err error) {
	resp_obj := proto.NewSendMsg("AUTH", "NOP")
	resp, err = user.Sess.Coder.Encode(resp_obj)
	return
}
