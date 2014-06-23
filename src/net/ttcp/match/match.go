package match

// 分配玩家到哪个桌的模块

import (
	"fmt"
)

import (
	"net/ttcp/types"
	"zmap"
)

var Desks *zmap.SafeMap

type Desk struct {

}

type Match struct {
	chan MQ string
}

func (m *Match) Start() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Caught panic in match goroutine")
			panic(x)
		}
	}()

	for {
		select {
		case data := <-m.MQ:
			fmt.Println(data)
	}
}

// 随机分配
// 调用这个函数要加锁
// TODO: 返回一个desk,在同一个桌里的user返回同一个desk
func (m *Match)RandAlloc(user *types.User) {

}

func init() {
	Desks = zmap.NewSafeMap()
}