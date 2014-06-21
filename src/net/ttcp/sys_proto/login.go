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
	resp_obj := proto.NewSendMsg("SYS", "LOGIN")
	if true {
		fmt.Println("Login:", username, password)
		// TODO: 登陆, 获得ID
		var uid uint32 = 88
		var user *types.User
		old_user_ := types.Users.Get(uid)
		if old_user_ != nil {

			user = old_user_.(*types.User)
			if user.IsActive { // 异地登陆,也是重连, 可以不在游戏中
				user.Disconnect() // 只需要关闭连接,旧session在main的Clear里会清理映射
				user.Sess = sess  // 此时为新user了
				fmt.Println("Clear old session, new user id:", user.Sess.ID)
			} else { // 重连, 一定是在游戏中,因为没在游戏中的断线都处理了
				user.Sess = sess
				user.IsActive = true
				fmt.Println("Reconnect session:", user.Sess.ID)
			}
		} else {
			user = types.NewUser(uid)
			types.Users.Set(uid, user)
		}
		user.Logined = true
		types.SessID2UID.Set(sess.ID, uid)
		(*resp_obj)["result"] = proto.R{Code: 0, Message: "ok"}
		resp = make([]byte, 100)
		resp, err = sess.Coder.Encode(resp_obj)

		fmt.Println("Logined")
	} else {
		err = errors.New("login failed")
	}

	return
}
