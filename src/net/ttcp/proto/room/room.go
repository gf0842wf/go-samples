package room

// 房间共用消息

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

// ..

// 开始游戏
func handle_gaming(user *types.User, obj *proto.Msg) (resp []byte, err error) {

	user.InGaming = true
	return
}
