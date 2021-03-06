package types

// 连接类型

import (
	"fmt"
	"net"
)

import (
	"net/ttcp/codec"
)

// TODO: 这里定义Session
/* Session应该包含的字段
IP net.IP
MQ chan IPCObj // Player's Internal Message Queue  (IPCObj包括发送/接受方ID, 消息(json string), 时间等)
Encoder
Decoder
User // User包含玩家基本信息(ID,昵称,所在游戏类型-游戏编号)(区别Player, Player包含User和游戏内部具体信息)
TODO: User里包含Session, Player里包含User
LoggedIn bool
KickOut bool
各种时间信息

-- MQ说明:
	MQ-客户端之间的消息: 解析收到的消息, 如果有MQ消息(比如玩家之间聊天消息)-->把消息投送到收件人们的sess.MQ中
		   -->收件人监听到自己的sess.MQ中有消息到来,然后解析消息
		   -->收件人通过outTag.Send把消息发给自己的客户端
	MQ-客户端之间的ping: 客户端A(ping)-->A:sess.inChs(ping)-->B:sess.MQ(ping)-->A:sess.MQ(pong)-->客户端A(pong)
*/
type Session struct {
	ID uint32
	IP net.IP

	Coder *codec.Coder

	Sender *SenderBuffer
}

func NewSession() *Session {
	SessID = (SessID + 1) % (2<<31 - 1)
	return &Session{ID: SessID}
}

// 强制断开连接
func (sess *Session) Disconnect() {
	fmt.Println("Disconnect:", sess.ID)
	sess.Sender.ctrl <- false
}
