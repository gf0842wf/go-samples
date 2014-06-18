package types

// 用户类型

import (
	"fmt"
)

type User struct {
	ID       int32
	Username string
	Nickname string
	Password string
	IsDummy  bool

	GameType int32  // 游戏类型,eg.斗地主
	GameRoom int32  // 在哪房间
	GameDesk int32  // 在哪桌
	GameID   uint64 // 游戏唯一编号

	Sess *Session
}

// User发送消息
func (user *User) Send(msg []byte) {
	err := user.Sess.Sender.Send(msg)
	if err != nil {
		fmt.Println("User.Send, Cannot send to client:", err)
	}
}

// 强制断开连接
func (user *User) Disconnect() {
	fmt.Println("Disconnect:", user.ID, user.Username)
	user.Sess.Sender.ctrl <- false
}

// 发送最后一个消息并断线
func (user *User) SendLose(msg []byte) {
	user.Send(msg)
	user.Disconnect()
}
