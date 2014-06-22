package auth

/*
心跳
<={"kind":"AUTH", "type":"NOP"}
=>{"kind":"AUTH", "type":"NOP"}

握手
<={"kind":"AUTH", "type":"PRESHAKE"} // 客户端发送握手准备
=>{"kind":"AUTH", "type":"REQSHAKE", "key":1234} // 服务器发送key给客户端
<={"kind":"AUTH", "type":"ACKSHAKE", "result":{"code":0, "message":'0k'}} // 客户端回应握手结果

登陆
<={"kind":"AUTH", "type":"LOGIN", user":'fk', "password":'112358'}
=>{"kind":"AUTH", "type":"LOGIN", "result":{"code":0, "message":'0k'}}

*/

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

var SysProtoHandlers map[string]func(*types.User, *proto.Msg) (resp []byte, err error)

func init() {
	SysProtoHandlers = map[string]func(*types.User, *proto.Msg) (resp []byte, err error){
		"AUTH.NOP":      handle_nop,
		"AUTH.PRESHAKE": handle_preshake,
		"AUTH.ACKSHAKE": handle_ackshake,
		"AUTH.LOGIN":    handle_login,
	}
}
