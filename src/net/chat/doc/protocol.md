# 聊天服务器协议 #

## 好友关系维护 ##
采用redis  

	采用set
	插入好友时,记得检查该好友是否存在在其它分组中
	group_i_uid:set(1,2,3,4,4,5) # i-第几个分组, uid-用户的id

	1.用户 36 添加好友 38 到分组 2
	is_exist = 0
	for i in count_groups:
		if `SISMEMBER group_i_36 38`:
			if i != 2:
				`SREM group_i_36 38`
			is_exist = 1
	`SADD group_2_36 38`

	2.用户 36 从分组 2 删除好友 38
	考虑到客户端可能不发送分组名字2,需要遍历来删除
	for i in count_groups:
		`SREM group_i_36 38` 
