package types

import (
	"zmap"
)

// 生成session, 在创建session时自增
var SessID uint32 = 0

// TODO: 这个map好像没用
// sessid 2 session
var Sessions *zmap.SafeMap = zmap.NewSafeMap()

// uid 2 user
var Users *zmap.SafeMap = zmap.NewSafeMap()

// sessid 2 uid
var SessID2UID *zmap.SafeMap = zmap.NewSafeMap()
