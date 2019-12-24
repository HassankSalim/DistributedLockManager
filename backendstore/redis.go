package backendstore

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisStore struct {
	client *redis.Client
}


func (rs *RedisStore) SetIfNotExists(key, token string, ttl int, ch chan bool) {
	success, err := rs.client.SetNX(key, token, time.Second * 60 * 5).Result()
	if err != nil {
		fmt.Printf("Err occured during getting key from redis: %s", err)
	}
	ch <- success
}

func (rs *RedisStore) DelIfKeyHasVal(key, token string, ch chan struct{}) {
	keyVal, err := rs.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("Key %s not found \n", key)
		ch <- struct{}{}
		return
	}
	if err != nil {
		fmt.Printf("Err occured during getting %s key \n", key)
		ch <- struct{}{}
		return
	}
	if keyVal == token {
		rs.client.Del(key)
	}
	ch <- struct{}{}
	fmt.Printf("deleted key %s \n", key)
	return
}

func NewRedisStore(addr string, db int) Store {
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr:addr,
			DB: db,
		}),
	}
}