package login

/*
心跳
<={"kind":"SYS", "type":"NOP"}
=>{"kind":"SYS", "type":"NOP"}

握手
<={"kind":"SYS", "type":"PRESHAKE"} // 客户端发送握手准备
=>{"kind":"SYS", "type":"REQSHAKE", "key":1234} // 服务器发送key给客户端
<={"kind":"SYS", "type":"ACKSHAKE", "result":{"code":0, "message":'0k'}} // 客户端回应握手结果

登陆
<={"kind":"SYS", "type":"LOGIN", user":'fk', "password":'112358'}
=>{"kind":"SYS", "type":"LOGIN", "result":{"code":0, "message":'0k'}}

*/

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

var SysProtoHandlers map[string]func(*types.User, *proto.Msg) (resp []byte, err error)

func init() {
	SysProtoHandlers = map[string]func(*types.User, *proto.Msg) (ack []byte, err error){
		"SYS.NOP":      handle_nop,
		"SYS.PRESHAKE": handle_preshake,
		"SYS.ACKSHAKE": handle_ackshake,
		"SYS.LOGIN":    handle_login,
	}
}
