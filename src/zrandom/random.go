package zrandom

// 随机数模块,使用时间种子

import (
	"math/rand"
	"time"
)

// 生成随机种子
func _Seed() {
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
// 算法的复杂度为O(N).
func Shuffle(pokers interface{}) {
	switch value := pokers.(type) {
	case []byte:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []int:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []uint16:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []uint32:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []uint64:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	case []string:
		size := len(value)
		for i := 0; i < size; i++ {
			x := Randint(i, size-1)
			value[i], value[x] = value[x], value[i]
		}
	default:

	}
}

// 摸牌, 从切片中随机选一张
// 注意这个函数的使用(因为返回的是interface{}):
// s := []byte{3, 4, 5}
// i := zrandom.Choice(s).(byte)
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
	case []bool:
		idx := Randint(0, len(value)-1)
		result = value[idx]
	default:
		result = nil
	}

	return
}

// 取样, 从切片中随机选n张
func Sample(pokers interface{}, n int) interface{} {
	switch value := pokers.(type) {
	case []byte:
		size := len(value)
		temp := make([]bool, size)
		idxs := make([]int, n)
		i := 0
		for i < n {
			// TODO: 这里如果运气不好, 会卡很长时间
			idx := Randint(0, size-1)
			if !temp[idx] {
				idxs[i] = idx
				temp[idx] = true
				i++
			}
		}
		smp := make([]byte, n)
		for k, v := range idxs {
			smp[k] = value[v]
		}
		return smp
	case []int:
		size := len(value)
		temp := make([]bool, size)
		idxs := make([]int, n)
		i := 0
		for i < n {
			// TODO: 这里如果运气不好, 会卡很长时间
			idx := Randint(0, size-1)
			if !temp[idx] {
				idxs[i] = idx
				temp[idx] = true
				i++
			}
		}
		smp := make([]int, n)
		for k, v := range idxs {
			smp[k] = value[v]
		}
		return smp
	case []uint16:

	case []uint32:
		size := len(value)
		temp := make([]bool, size)
		idxs := make([]int, n)
		i := 0
		for i < n {
			// TODO: 这里如果运气不好, 会卡很长时间
			idx := Randint(0, size-1)
			if !temp[idx] {
				idxs[i] = idx
				temp[idx] = true
				i++
			}
		}
		smp := make([]uint32, n)
		for k, v := range idxs {
			smp[k] = value[v]
		}
		return smp
	case []uint64:

	case string:

	case []string:
		size := len(value)
		temp := make([]bool, size)
		idxs := make([]int, 0)
		i := 0
		for i < n {
			// TODO: 这里如果运气不好, 会卡很长时间
			idx := Randint(0, size-1)
			if !temp[idx] {
				idxs = append(idxs, idx)
				temp[idx] = true
				i++
			}
		}
		smp := make([]string, n)
		for k, v := range idxs {
			smp[k] = value[v]
		}
		return smp
	default:

	}
	return nil
}
