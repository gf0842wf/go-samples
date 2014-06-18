package types

import (
	"zmap"
)

// 生成session, 在创建session时自增
var SessID uint64 = 0

// sessid 2 session
var Sessions *zmap.SafeMap = zmap.NewSafeMap()

// uid 2 user
var Users *zmap.SafeMap = zmap.NewSafeMap()
