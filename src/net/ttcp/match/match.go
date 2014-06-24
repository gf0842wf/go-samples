package match

// 分配玩家到哪个桌的模块

import (
	"errors"
	"fmt"
	"time"
)

import (
	"net/ttcp/types"
	// "zrandom"
)

// 这个Desk类可以继承
type Desk struct {
	DeskNo    int
	Status    int // 0-未开始, 1-游戏中, 2-结束游戏
	GameID    uint32
	StartTime time.Time
	EndTime   time.Time

	MQ    chan string
	Users [MAX_MEMBERS]*types.User // 数组很小,用值拷贝没问题
}

func (d *Desk) Add(user *types.User) (r, event bool) {
	// r表示Add是否成功, event是事件是否到来
	for i, v := range d.Users {
		if v == nil {
			d.Users[i] = user
			r = true
			break
		}
	}
	// 检查人满事件,如果满了并且只有一个人没有准备,开始倒计时踢那个人
	if d.EptPos() == 0 {
		// 事件到来,人满了
		event = true
		// TODO:
	}

	return
}

func (d *Desk) Remove(user *types.User) (r, event bool) {
	for i, v := range d.Users {
		if v == user {
			d.Users[i] = nil
			r = true
			break
		}
	}
	// 检查空桌事件,如果房间没人,爆房,关闭goroutine
	if d.EptPos() == MAX_MEMBERS {
		// 事件到来,桌空了
		event = true
		// TODO:
	}

	return
}

func (d *Desk) Delete(uid uint32) (bool, bool) {
	user_ := types.Users.Get(uid)
	if user_ != nil {
		return d.Remove(user_.(*types.User))
	}
	return false, false
}

func (d *Desk) PrepUser(user *types.User) (r, event bool) {
	for _, v := range d.Users {
		if v == user {
			user.IsPrep = true
			r = true
			break
		}
	}

	// 检查游戏开始事件(全部准备/某些游戏MIN_MEMBERS人准备就可以玩),开始游戏循环
	cnt := 0
	for _, u := range d.Users {
		if u.IsPrep {
			cnt++
		}
	}
	if cnt >= MIN_MEMBERS {
		// 事件到来,可以开始游戏了
		event = true
		// TODO:
	}

	return
}

// 空闲座位数
func (d *Desk) EptPos() (cnt int) {
	cnt = 0
	for _, v := range d.Users {
		if v == nil {
			cnt++
		}
	}
	return
}

// 更新users游戏开始状态, deskno等
func (d *Desk) UpdateUserGamingInfo() (err error) {
	for _, u := range d.Users {
		if u == nil {
			err = errors.New("not full")
			return
		} else if !u.IsPrep {
			err = errors.New("some one not prep")
			return
		}
		u.InGaming = true
		u.DeskNo = d.DeskNo
	}
	return
}

func (d *Desk) Start() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Caught panic in desk goroutine")
			panic(x)
		}
	}()

	d.UpdateUserGamingInfo()
	d.Status = 1

	// 主游戏循环,游戏逻辑写在这里
	for {
		select {
		case data := <-d.MQ:
			fmt.Println(data)
		}
	}
}

func NewDesk() *Desk {
	ret := &Desk{}
	for i := 0; i < MAX_MEMBERS; i++ {
		ret.Users[i] = nil
	}
	return ret
}

// ---------------------------------------

type DSManage [DESK_SIZE]*Desk

func (ds DSManage) EptPoses() (poses []int) {
	poses = make([]int, len(ds))
	for i, d := range ds {
		poses[i] = d.EptPos()
	}
	return
}

func (ds DSManage) IndexUser(user *types.User) (deskno int, err error) {
	for i, d := range ds {
		for _, u := range d.Users {
			if user == u {
				deskno = i
				return
			}
		}
	}
	err = errors.New("user not found")
	return
}

func (ds DSManage) FindUser(user *types.User) (desk *Desk, err error) {
	deskno, e := ds.IndexUser(user)
	if e != nil {
		err = e
		return
	}
	desk = ds[deskno]
	return
}

// 存放所有desk
var Desks DSManage

// -------------------------------------------------------

// 随机分配
// 消息中调用
// 调用这个函数要加锁
func RandAlloc(user *types.User) (err error) {
	// poses := Desks.EptPoses() // poses [4, 2, 4, 0, MAX_MEMBERS, ...]
	// probs := make([]int, len(poses))
	// for i, ept := range poses {
	// 	if ept == 0 {
	// 		probs[i] = 0
	// 	} else {
	// 		probs[i] = MAX_MEMBERS - ept
	// 	}
	// }
	// deskno, _ := zrandom.ProbChoiceI(probs, poses)

	poses := Desks.EptPoses() // poses [4, 2, 4, 0, MAX_MEMBERS, ...]
	// 找到空位数最小且桌号最小的桌(求空位的最小值且是从后往前比较)
	deskno := 0
	for i := len(poses) - 1; i > 0; i-- {
		tmp := poses[i]
		if !HALFWAY_ENTER && Desks[i].Status == 1 { // 游戏不允许中途进入
			tmp = 0
		}
		if poses[deskno] >= tmp && tmp != 0 {
			deskno = i
		}
	}
	desk := Desks[deskno]
	if poses[deskno] == 0 {
		err = errors.New("no empty desk")
		return
	} else if !HALFWAY_ENTER && desk.Status == 1 {
		err = errors.New("gaming not allow enter")
		return
	}
	desk.Add(user)
	return
}

// 固定分配
// 消息中调用
// 调用这个函数要加锁
func FixAlloc(user *types.User, deskno int) (err error) {
	desk := Desks[deskno]
	if desk.EptPos() == 0 {
		err = errors.New("the desk is full")
		return
	}
	desk.Add(user)
	return
}

// 玩家准备
// 消息中调用
func UserPrep(user *types.User) (err error) {
	// 从Desks中找到这个user所在的desk, 告诉它user准备了,desk检查一下所有用户是否准备,然后确认这个桌是否可以开始
	desk, e := Desks.FindUser(user)
	if e != nil {
		err = e
		return
	}
	desk.PrepUser(user)
	return
}

// 玩家准备
// 消息中调用/掉线调用
func UserLeave(user *types.User, action int) (err error) {
	// action-0意外断线等, action-1暴力离开
	desk, e := Desks.FindUser(user)
	if e != nil {
		err = e
		return
	}
	user.IsPrep = false
	if desk.Status == 1 { // 游戏中
		if action == 1 { // 玩家执意要离开(逃跑)
			desk.Remove(user)
			user.InGaming = false
			user.DeskNo = -1
		} else {
			// 意外断线离开不清理玩家
		}
	} else { // 没在游戏中, 马上清除,在main的Clear中调用 UserLeave(user, 0)
		desk.Remove(user)
		user.InGaming = false
		user.DeskNo = -1
	}
	return
}

func init() {
	for i := 0; i < len(Desks); i++ {
		Desks[i] = NewDesk()
		Desks[i].DeskNo = i
	}
}
