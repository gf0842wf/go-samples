package proto

import (
	"encoding/json"
	"strconv"
)

// 收发的消息
type Msg map[string]interface{}

// kind and type 判断
func (msg *Msg) EqualKT(k, t string) (r bool, err error) {
	if k_, ok := msg["kind"]; ok {
		if t_, ok := msg["type"]; ok {
			r := (k == k_.(string) && t == t_.(string))
		} else {
			err = error.Error("not type field")
		}
	} else {
		err = error.Error("not kind field")
	}
	return
}

func (msg *Msg) Json() []byte {
	jsobj, _ := json.Marshal(msg)
	return jsobj
}

func (msg *Msg) Repr() string {
	repr := string(msg.Json())
	return jsobj
}

// result 字段
type R struct {
	Code    int    `code`
	Message string `message`
}

func (r *R) Repr() string {
	repr := string(r.Json())
	return repr
}

func (r *R) Json() []byte {
	jsobj, _ := json.Marshal(r)
	return jsobj
}
