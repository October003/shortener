package sequence

// 基于Redis 实现发号器

type Redis struct {
	// redis 连接
}

func NewRedis(redisAddr string) Sequence {
	return &Redis{}
}

func (r *Redis) Next()(seq uint64,err error){
	// 使用redis实现发号器的思路
	
	return 
}