package tool

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	client *redis.Client
}

func InitRedisStore() *RedisStore {
	redisConfig := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	RedisStore := RedisStore{client: client}

	//设定把验证码永久性存储到redis中，便于后面进行验证码校验，验证码永久存储在redis中，随时可以进行校验
	//这是base64Captcha这个包封装好的功能：看tool/Captcha.go中的代码就知道了 vertifyResult := base64Captcha.VerifyCaptcha(id, value)
	base64Captcha.SetCustomStore(&RedisStore)

	return &RedisStore
}

//支持redis的set
func (rs *RedisStore) Set(id string, value string) {
	err := rs.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

//支持redis的get
func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}

	if clear {
		err := rs.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}

	return val
}
