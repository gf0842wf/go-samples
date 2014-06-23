package match

// 分配玩家到哪个桌的模块

import (
	"fmt"
	"time"
)

import (
	"net/ttcp/types"
	"zmap"
)

// 这个Desk类可以继承
type Desk struct {
	Status    int // 0-未开设, 1-游戏中, 2-结束游戏
	GameID    uint32
	StartTime time.Time
	EndTime   time.Time

	MQ    chan string
	Users [MAX_MEMBERS]*types.User // 数组很小,用值拷贝没问题
}

func (d *Desk) Add(user *types.User) bool {
	for i, v := range d.Users {
		if v == nil {
			d.Users[i] = user
			return true
		}
	}
	return false
}

func (d *Dsek) Remove(user *types.User) bool {
	for i, v := range d.Users {
		if v == user {
			d.Users[i] = nil
			return true
		}
	}
	return false
}

func (d *Desk) Delete(uid uint32) bool {
	user_ := types.Users.Get(uid)
	if user_ != nil {
		return d.Remove(user_.(*types.User))
	}
	return false
}

// 空闲座位数
func (d *Desk) EptPos() (cnt int) {
	cnt = 0
	for i, v := range d.Users {
		if v == nil {
			cnt++
		}
	}
	return
}

func NewDesk() *Desk {
	ret = &Desk{}
	for i := 0; i < MAX_MEMBERS; i++ {
		ret.Users[i] = nil
	}
	return ret
}

var Desks [DESK_SIZE]*Desk

func (d *Desk) Start() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Caught panic in desk goroutine")
			panic(x)
		}
	}()

	for {
		select {
		case data := <-d.MQ:
			fmt.Println(data)
		}
	}
}

// 随机分配
// 调用这个函数要加锁
func RandAlloc(user *types.User) {
	for num, desk := range Desks {

	}
}

func init() {

}
