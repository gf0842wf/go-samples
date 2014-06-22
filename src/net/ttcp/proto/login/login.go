package login

// 登陆消息

import (
	"errors"
	"fmt"
)

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

func handle_login(user *types.User, obj *proto.Msg) (resp []byte, err error) {
	if !user.Sess.Coder.Shaked {
		err = errors.New("not shaked")
		return
	}
	username, ok := (*obj)["username"]
	if !ok {
		err = errors.New("no username field")
		return
	}
	password, ok := (*obj)["password"]
	if !ok {
		err = errors.New("no password field")
		return
	}
	if true {
		fmt.Println("Login:", username.(string), password.(string))
		// TODO: 登陆, 获得ID
		var uid uint32 = 88
		old_user_ := types.Users.Get(uid)
		if old_user_ != nil {
			old_user := old_user_.(*types.User)
			if old_user.IsActive { // 异地登陆,也是重连, 可以不在游戏中
				old_user.Disconnect()     // 关闭旧sess
				old_user.Sess = user.Sess // 只保留old_user的user信息,不要sess信息
				fmt.Println("Clear old session, new user id:", old_user.Sess.ID)
			} else { // 重连, 一定是在游戏中,因为没在游戏中的断线都处理了
				old_user.Sess = user.Sess
				old_user.IsActive = true
				fmt.Println("Reconnect session:", user.Sess.ID)
			}
		} else { // 新登陆
			user.ID = uid
			types.Users.Set(uid, user)
		}
		user.Logined = true
		// 回应登陆成功消息
		resp_obj := proto.NewSendMsg("SYS", "LOGIN")
		(*resp_obj)["result"] = proto.R{Code: 0, Message: "ok"}
		resp, err = user.Sess.Coder.Encode(resp_obj)

		fmt.Println("Logined")
	} else {
		err = errors.New("login failed")
	}

	return
}
