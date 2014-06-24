package types

// 用户类型

import (
	"fmt"
)

type User struct {
	ID       uint32
	Username string
	Nickname string
	Password string

	IsDummy  bool
	KickOut  bool // 被踢标志
	IsActive bool // 是否在线
	IsPrep   bool // 是否准备
	InGaming bool // 是否游戏中
	Logined  bool // 是否登陆

	GameType int32  // 游戏类型,eg.100:斗地主
	RoomID   int32  // 在哪房间, RoomID映射到RoomInfo
	DeskNo   int    // 所在桌号,-1表示没有游戏中
	GameID   uint32 // 正在玩的游戏的唯一编号(自增)

	MQ chan IPCObj // User之间通信队列

	Sess *Session
}

func NewUser() *User {
	return &User{IsActive: true}
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
	user.Sess.Disconnect()
}

// 发送最后一个消息并断线
func (user *User) SendLose(msg []byte) {
	user.Send(msg)
	user.Disconnect()
}
