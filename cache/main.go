package cache

import (
	//"strconv"
	"zixu.com/cmdAdmin/conf/vars"
	"zixu.com/cmdAdmin/util"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	//db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	r := vars.Redis
	client := redis.NewClient(&redis.Options{
		Addr:       r.Addr,
		Password:   r.Password,
		DB:         r.DB,
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
