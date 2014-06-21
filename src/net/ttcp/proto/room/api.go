package room

/*
进入房间(普通)
<={"kind":"ROOM", "type":"INTOROOM", "subtype":0, "gametype":100, "roomid":1001} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"INTOROOM", "result":{"code":0, "message":'0k'}}
进入房间并进入桌(百人)
<={"kind":"ROOM", "type":"INTOROOM", "subtype":1, "gametype":100, "roomid":1001} subtype:0-普通,1-百人
=>{"kind":"ROOM", "type":"INTOROOM", "result":{"code":0, "message":'0k'}}
快速加入桌(仅普通)
<={"kind":"ROOM", "type":"QUICKDESK"}
=>{"kind":"ROOM", "type":"QUICKDESK", "result":{"code":0, "message":'0k'}}
选择桌加入(仅普通)
<={"kind":"ROOM", "type":"INTODESK", "deskno":1}
=>{"kind":"ROOM", "type":"INTODESK", "result":{"code":0, "message":'0k'}}
以下协议根据情况而定
*/
