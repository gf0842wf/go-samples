package sys_proto

// 握手信息

import (
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_preshake(sess *types.Session, obj *proto.Msg) (ack []byte, err error) {
	ack_obj := proto.NewSendMsg("SYS", "REQSHAKE")
	(*ack_obj)["key"] = sess.Coder.CryptKey
	ack = make([]byte, 60)
	sess.Coder.Encode(ack_obj, ack)
	return
}

func handle_ackshake(sess *types.Session, obj *proto.Msg) (ack []byte, err error) {
	sess.Coder.Shaked = true
	fmt.Println("Shaked:", sess.ID)
	return
}
