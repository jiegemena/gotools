package myredis

import (
	"context"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

type Myredis struct {
	rdb *redis.Client
}

func (r *Myredis) GetRedisClient(addr string, pwd string, db int) {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})
	// r.rdb = rdb
}

func (r *Myredis) Set(key string, val string, timeout int) error {

	err := r.rdb.Set(ctx, key, val, 0).Err()
	// err := s.Err()
	return err
}

func (r *Myredis) Get(key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	return val, err
}

func (r *Myredis) Close(key string) {
	r.rdb.Close()
}
