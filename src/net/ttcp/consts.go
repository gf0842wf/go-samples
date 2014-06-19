package main

// main中用的常量

const (
	ADDR                 = ":8888"
	TCP_TIMEOUT          = 120
	MAX_DELAY_IN         = 120
	DEFAULT_INQUEUE_SIZE = 16  // 队列的size, 不是数据大小,是MAX_RECV_DATA_SIZE这么大的数据的16维数组
	MAX_RECV_DATA_SIZE   = 512 // 接受消息缓存大小(不含头)
)
