# 聊天服务器协议 #

## 前端部署 ##
haproxy进行tcp负载均衡反向代理

## 好友关系维护 ##
采用redis  

	采用set # TODO: 用sorted set http://blog.csdn.net/suiye/article/details/7765329
	分组用户好友关系(set):group_i_uid:set(1,2,3,4,4,5) # i-group id, uid-用户的id
    用户分组映射(hash):groups_uid: 

	1.用户 36 添加好友 38 到分组 2
	is_exist = 0
	for i in group_ids:
		if `SISMEMBER group_i_36 38`:
			if i != 2:
				`SREM group_i_36 38`
			is_exist = 1
	`SADD group_2_36 38`

	2.用户 36 (从分组 2) 删除好友 38
	考虑到客户端可能不发送分组名字2,需要遍历来删除
	for i in group_ids:
		`SREM group_i_36 38` 

    3.获取用户 36 的所有好友
    friends = []
    for i in group_ids:
        friend = {}
        friend["gid"] = i # gid-group id
        friend[members] = `SMEMBERS group_i_36`
        friends.append(friend)

    


