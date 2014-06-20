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
	IsActive bool // 连接是否关闭
	InGaming bool // 是否游戏中
	Logined  bool // 是否登陆

	GameType int32  // 游戏类型,eg.斗地主
	GameRoom int32  // 在哪房间
	GameDesk int32  // 在哪桌
	GameID   uint64 // 游戏唯一编号

	Sess *Session
}

func NewUser(userID uint32) *User {
	return &User{ID: userID, IsActive: true}
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
