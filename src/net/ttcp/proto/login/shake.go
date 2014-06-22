package login

// 握手信息

import (
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_preshake(user *types.User, obj *proto.Msg) (ack []byte, err error) {
	ack_obj := proto.NewSendMsg("SYS", "REQSHAKE")
	(*ack_obj)["key"] = user.Sess.Coder.CryptKey
	ack, err = user.Sess.Coder.Encode(ack_obj)
	return
}

func handle_ackshake(user *types.User, obj *proto.Msg) (ack []byte, err error) {
	user.Sess.Coder.Shaked = true
	fmt.Println("Shaked:", user.Sess.ID)
	return
}
