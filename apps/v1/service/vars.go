package service

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "sync"
    "unity.service/pkg/org.unity.sdk/setting"

    //"github.com/hibiken/asynq"

)
var o  sync.Once
var c  *redis.Client
type Edge struct {
    Ctx *context.Context
    //AsyncRedis *asynq.Client
    Redis *redis.Client
    L *sync.Mutex
    C *sync.Cond
    M *sync.RWMutex
    G *sync.WaitGroup
}


func Initialization() *Edge {
    return &Edge{
        Ctx: nil,
        //AsyncRedis: asynq.NewClient(asynq.RedisClientOpt{Addr: setting.RedisConf.Host, Password: setting.RedisConf.Password, DB: 10}),
        Redis: Rs(),
        L: new(sync.Mutex),
        C: new(sync.Cond),
        M: new(sync.RWMutex),
    }

}

func Rs() *redis.Client {

    if c != nil {
        return c
    } else {
        o.Do(func() {
            c = redis.NewClient(&redis.Options{Addr: setting.RedisConf.Host, Password:setting.RedisConf.Password, DB: setting.RedisConf.Db, ReadTimeout:300e6, WriteTimeout: 300e6})
        })
    }
   return c

}


func (this Edge) RKey (key string, part string) string {
    return fmt.Sprintf(key, part)
}


func (this Edge) Chunk(length, count int, fc func(offset, size int)) {
    for i := 0; i < length/count; i++ {
        fc(i*count, (i+1)*count)
    }
}


const (
    EnvStrRedisQrcodeWithValidTime = `iss:qrcode:%s`
    EnvStrRedisGpsFlag = `iss:qrcode:gps:%s`
)



