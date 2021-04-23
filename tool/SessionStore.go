package tool

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//session使用方式：基于cookie,基于redis，基于memcached

//初始化session操作
func InitSession(engine *gin.Engine) {
	//获取redis配置文件内容
	redisConfig := GetConfig().RedisConfig
	//redis新建连接
	store, err := redis.NewStore(10, "tcp", redisConfig.Addr+":"+redisConfig.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//调用session写入redis
	engine.Use(sessions.Sessions("mysession", store))
}

//set操作session
func SetSession(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}

	session.Set(key, value)

	return session.Save()
}

//get操作session
func GetSession(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session
}
