package login

// 心跳

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_nop(sess *types.Session, obj *proto.Msg) (resp []byte, err error) {
	resp_obj := proto.NewSendMsg("SYS", "NOP")
	resp, err = sess.Coder.Encode(resp_obj)
	return
}
