package zrandom

// 随机数模块,使用时间种子

import (
	"math/rand"
	"time"
)

func _Seed() {
	// 生成随机种子
	rand.Seed(time.Now().UTC().UnixNano())
}

// 产生 min <= N<= max 之间的一个随机数整数
func Randint(min, max int) int {
	if max-min <= 0 {
		return min
	}
	_Seed()
	return min + rand.Intn(max-min+1)
}

// 洗牌, 把切片打乱, inplace 操作
// 1.随机产生一个1-n的数x,然后让第x张牌和第1张牌互相调换.
// 2.随机产生一个1-n的数y,然后让第y张牌和第2张牌互相调换.
// 3.随机产生一个1-n的数z,然后让第z张牌和第i张牌互相调换.(i=3,4,5...54)
// 算法的复杂度为O(N).
func Shuffle(pokers interface{}) {
	switch value := pokers.(type) {
	case []byte:
		size := len(value)
		for i := 1; i < size; i++ {
			x := Randint(0, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []int:

	case []uint16:

	case []uint32:

	case []uint64:

	case string:

	case []string:

	default:

	}
}

// 摸牌, 从切片中随机选一张
func Choice(pokers interface{}) (result interface{}) {
	// if value, ok := pokers.([]byte); ok {
	// 	idx := Randint(0, len(value)-1)
	// 	val = value[idx]
	// }
	switch value := pokers.(type) {
	case []byte:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case []int:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case []uint16:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case []uint32:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case []uint64:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case string:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	case []string:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	default:
		result = nil
	}

	return
}

// 取样, 从切片中随机选n张
func Sample(pokers interface{}, n int) (smp interface{}) {
	// switch value := pokers.(type) {
	// case []byte:

	// case []int:

	// case []uint16:

	// case []uint32:

	// case []uint64:

	// case string:

	// case []string:

	// default:

	// }
	return
}
