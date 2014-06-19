package sys_proto

// 登陆消息

import (
	"errors"
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_login(sess *types.Session, obj *proto.Msg) (resp []byte, err error) {
	username, ok := obj["username"]
	if !ok {
		return
	}
	password, ok := obj["password"]
	if !ok {
		return
	}
	resp_obj := proto.NewSendMsg("SYS", "LOGIN")
	if true { // TODO: 登陆, 获得ID
		fmt.Println("Login:", username, password)
		id := 88
		old_id := types.Users.Get(id)
		if old_id != nil {
			// 重连
		}
		(*resp_obj)["result"] = proto.R{Code: 0, Message: "ok"}
		resp = make([]byte, 100)
		sess.Coder.Encode(resp_obj, resp)

		fmt.Println("Logined")
	} else {
		err = errors.New("login failed")
	}

	return
}
