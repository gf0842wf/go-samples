package proto

import (
	"encoding/json"
	"errors"
)

// ----------收发的消息-------------
type Msg map[string]interface{}

// kind and type 判断
func (msg Msg) EqualKT(k, t string) (r bool, err error) {
	if k_, ok := msg["kind"]; ok {
		if t_, ok := msg["type"]; ok {
			r = (k == k_.(string) && t == t_.(string))
		} else {
			err = errors.New("not type field")
		}
	} else {
		err = errors.New("not kind field")
	}
	return
}

// 返回kind
func (msg Msg) K() (k string, err error) {
	k_, ok := msg["kind"]
	k = k_.(string)
	if !ok {
		err = errors.New("not kind field")
	}
	return
}

// 返回type
func (msg Msg) T() (t string, err error) {
	t_, ok := msg["type"]
	t = t_.(string)
	if !ok {
		err = errors.New("not type field")
	}

	return t, err
}

//返回kind.type
func (msg Msg) KT() (kt string, err error) {
	if k_, ok := msg["kind"]; ok {
		if t_, ok := msg["type"]; ok {
			kt = k_.(string) + "." + t_.(string)
		} else {
			err = errors.New("not type field")
		}
	} else {
		err = errors.New("not kind field")
	}
	return
}

func (msg *Msg) Json() (jsstr []byte, err error) {
	jsstr, err = json.Marshal(msg)
	return
}

func (msg *Msg) Repr() string {
	repr_, _ := msg.Json()
	return string(repr_)
}

func NewSendMsg(k, t string) *Msg {
	return &Msg{"kind": k, "type": t}
}

// --------------result 字段-----------
type R struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *R) Json() (jsstr []byte, err error) {
	jsstr, err = json.Marshal(r)
	return
}

func (r *R) Repr() string {
	repr_, _ := r.Json()
	return string(repr_)
}
