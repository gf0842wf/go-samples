# 聊天服务器协议 #

## 前端部署 ##
haproxy进行tcp负载均衡反向代理

	<haproxy.cfg>
	# this config needs haproxy-1.1.28 or haproxy-1.2.1
	
	global
		log 127.0.0.1	local0
		log 127.0.0.1	local1 notice
		#log loghost	local0 info
		maxconn 65535
		ulimit-n 131086
		#chroot /usr/share/haproxy
		user haproxy
		group haproxy
		daemon
		nbproc  5 # 五个并发进程
		pidfile /var/run/haproxy.pid
		#debug
		#quiet
	
	defaults
		#log	global
		mode	http
		option	httplog
		option	dontlognull
		retries	2
		option redispatch
		maxconn	4096
		contimeout	5000
		clitimeout	50000
		srvtimeout	50000
	
	########统计页面配置########
	listen admin_stats
		bind 0.0.0.0:9100               #监听端口 
		mode http                       #http的7层模式  
		option httplog                  #采用http日志格式 
		log 127.0.0.1 local0 err
		maxconn 10
		stats enable
		stats refresh 30s               #统计页面自动刷新时间  
		stats uri /                     #统计页面url  
		stats realm XingCloud\ Haproxy  #统计页面密码框上提示文本  
		stats auth admin:admin          #统计页面用户名和密码设置  
		stats hide-version              #隐藏统计页面上HAProxy的版本信息  
	
	########chat服务器配置############# 
	listen chat
		bind 0.0.0.0:9000
		mode tcp
		maxconn 100000
		log 127.0.0.1 local0 debug
		server s1 192.168.1.111:9001 weight 1 # 这个可以部署haproxy,但是最好别
		server s1 192.168.1.112:9001 weight 5
		server s1 192.168.1.113:9001 weight 5
		server s1 192.168.1.114:9001 weight 5
	########frontend配置############### 
`sudo haproxy -f /etc/haproxy/haproxy.cfg`


## 好友关系 ##
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
	#TODO 聊天室功能

## 状态 ##
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

	# TODO: 其他状态统计信息

## 离线消息 ##
采用redis

	用户到xx的离线消息(hash) msg:offline:`uid`=>`target_uid`=>`message`
	target_uid:见下面的消息协议说明
	message:
	{
		"content":"",
		"subject":0,
		"send_time":0,
	}

## 在线消息 ##
暂不保存

## 消息协议 ##
暂时采用msgpack  
中文采用utf8编码  
消息长度(4bytes)+消息信息+消息体

消息信息(4+1+8+8+8+4+1)

	{
		"type":0, # 0-查询消息, 1-聊天消息 (uint32)
		"online":0, # 0-离线消息, 1-在线消息 (byte)
		"from":36, # uid    uint64, [1, 50)系统预留, [50, 150)聊天室, [150, 1000)预留, [1000, +∞)用户
		"to":38, # target_uid    uint64, [1, 50)系统预留, [50, 150)聊天室, [150, 1000)预留, [1000, +∞)用户
		"send_time":0, # unix时间戳(单位ms, uint64)
		"flag1":uint32,
		"flag2":byte
	}

消息体

	string

