package types

// 全局变量, sess和user的映射

import (
	"zmap"
)

// 生成session, 在创建session时自增
var SessID uint32 = 0

// 生成GameID, 在正在游戏中时自增
var GameID uint32 = 0

// TODO: 这个map好像没用
// sessid 2 session
var Sessions *zmap.SafeMap = zmap.NewSafeMap()

// uid 2 user
var Users *zmap.SafeMap = zmap.NewSafeMap()

// sessid 2 uid
var SessID2UID *zmap.SafeMap = zmap.NewSafeMap()

// TODO: 增加一个map <UserInfos> uid 2 userinfo 的映射, user具体信息用的不多,所以放在映射里,不放在user字段里
// TODO: 增加一个map <RoomInfoS> roomid 2 roominfo 的映射
