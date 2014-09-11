# 聊天服务器协议 #

## 前端部署 ##
haproxy进行tcp负载均衡反向代理

## 好友关系维护 ##
采用redis  

	分组用户好友关系(set) relation:`gid`:`uid`=>set("1","2","3","4","4","5") # gid:group id, uid:用户的id
    用户分组(set) relation:gids:`uid`=>set("0","1","2","同事") 
	注: 分组名不能出现:号

	1.用户 36 添加分组'同事'  <'0'是默认分组>
	is_exist = 0
	if `SISMEMBER relation:gids:36 同事`:
		is_exist = 1
	else:
		`SADD relation:gids:36 同事`

	2.用户 36 删除分组'同事'  <'0'是默认分组>
	is_exist = 1
	if not `SISMEMBER relation:gids:36 同事`:
		is_exist = 0
	else:
		# 删除分组把好友移动到分组0
		`SUNIONSTORE relation:0:36 relation:0:36 relation:同事:36`
		`DEL relation:同事:36`
		`SREM relation:gids:36 同事`

	3.用户 36 添加好友 38 到分组 '同事'
	gid = '同事' or '0'
	gids = `SMEMBERS relation:gids:36`
	is_exist = 0
	for i in `gids`:
		if `SISMEMBER relation:`gid`:36 38`:
			if i != gid:
				`SREM relation:`gid`:36 38`
			is_exist = 1
	`SADD relation:`gid`:36 38`

	4.用户 36 (从分组 2) 删除好友 38
	考虑到客户端可能不发送分组名字2,需要遍历来删除
	for gid in `gid`:
		`SREM relation:`gid`:36 38` 

    5.获取用户 36 的所有好友
    groups = []
    for gid in `gids`:
        group = {}
        group["gid"] = `gid` # gid-group id
        group["members"] = `SMEMBERS relation:`gid`:36`
        groups.append(group)

	#TODO: 用户1和2的共同好友,用户1的好友数等等    
	#TODO 好友采用zset存储,score存储添加好友的时间戳?


## 状态维护 ##
采用redis

	用户所在服务器(kv) status:`uid`=>`sid` # sid:server id
	服务器所有用户(zset) status:`sid`=>uid=>time.time() # uid作为member, time.time()作为score(建立连接时的时间戳)
	
	1.用户36在服务器2上线
	`SET status:36 2`
	`ZADD status:2 time.time() 36`

	2.用户36所在服务器(是否在线)
	`GET status:36`

	3.用户36下线
	sid = `GET status:36`
	`ZREM status:`sid` 36`

	4.服务器2的用户数
	`ZCARD status:2`

	# TODO: 状态统计信息
