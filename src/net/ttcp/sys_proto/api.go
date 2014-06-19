package user_proto

/*
心跳
<={"kind":"SYS", "type":"NOP"}
-

握手
=>{"kind":"SYS", "type":"SHAKE", "key":1234}
<={"kind":"SYS", "type":"SHAKE", "result":{"code":0, "message":'0k'}}

登陆
<={"kind":"SYS", "type":"LOGIN", user":'fk', "password":'112358'}
=>{"kind":"SYS", "type":"LOGIN", "result":{"code":0, "message":'0k'}}

*/

var REQ_types = map[string]int16{
	"NOP":   0,    // 心跳包, 服务端不回应
	"SHAKE": 1,    // **客户端握手回应
	"LOGIN": 2,    // 登陆请求
	"TALK":  1000, // 文字消息
}

var RESP_types = map[string]int16{
	"NOP":   100,  // 心跳包, 服务端不回应
	"SHAKE": 101,  // **服务端握手请求
	"LOGIN": 202,  // 登陆回应
	"TALK":  1100, // 文字消息
}

var RREQ_types map[int16]string
var RRESP_types map[int16]string

var ProtoHandler map[int16]func(*Session, *packet.Packet) (resp []byte, err error)

func init() {
	for k, v := range REQ_types {
		RREQ_types[v] = k
	}
	for k, v := range RESP_types {
		RRESP_types[v] = k
	}
}
