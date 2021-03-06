package room

/*
用户加入/退出时 广播消息, 根据情况广播给房间每个人/牌桌每个人
=>{{"kind":"ROOM", "type":"BCROOM", "subtype":"IN", "uid":1123}
=>{{"kind":"ROOM", "type":"BCROOM", "subtype":"OUT", "uid":1123}
=>{{"kind":"ROOM", "type":"BCDESK", "subtype":"IN", "uid":1123}
=>{{"kind":"ROOM", "type":"BCDESK", "subtype":"OUT", "uid":1123}

进入房间(普通)
<={"kind":"ROOM", "type":"INTOROOM", "subtype":0, "gametype":100, "roomid":1001} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"INTOROOM", "result":{"code":0, "message":'0k'}, TODO: 房间信息}
进入房间并进入桌(百人)
<={"kind":"ROOM", "type":"INTOROOM", "subtype":1, "gametype":100, "roomid":1001} subtype:0-普通,1-百人
<={"kind":"MATCH", "type":"NORMAL", subtype:1} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"INTOROOM", "result":{"code":0, "message":'0k'}, TODO: 房间信息}
快速加入桌(仅普通)
<={"kind":"ROOM", "type":"QUICKDESK"}
<={"kind":"MATCH", "type":"QUICK", subtype:0} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"QUICKDESK", "result":{"code":0, "message":'0k'}, TODO: 牌桌信息}
选择桌加入(仅普通)
<={"kind":"ROOM", "type":"INTODESK", "deskno":1}
<={"kind":"MATCH", "type":"NORMAL", "deskno":1, subtype:0} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"INTODESK", "result":{"code":0, "message":'0k'}, TODO: 牌桌信息
在加入桌和游戏开始之间协议根据游戏不同而定
游戏开始
=>{"kind":"ROOM", "type":"GAMING"}
=>{"kind":"ROOM", "type":"GAMING", "result":{"code":0, "message":'0k'}}

继续:
发ROOM消息和MATCH消息
换桌:
发ROOM消息和MATCH消息
退出:
发ROOM消息和MATCH消息
重进:
发ROOM消息和MATCH消息
*/

import (
	"net/ttcp/proto"
	"net/ttcp/types"
)

var RoomProtoHandlers map[string]func(*types.User, *proto.Msg) (resp []byte, err error)

func init() {
	RoomProtoHandlers = map[string]func(*types.User, *proto.Msg) (resp []byte, err error){
		// ..
		"ROOM.GAMING": handle_gaming,
	}
}
