package redis

import (
	"fiber/helpers"
	"github.com/go-redis/redis"
)

// @see https://github.com/go-redis/redis
func ConnectDb(db int) *redis.Client{
	host := helpers.LoadConfiguration("redis", "redis_host")
	pass := helpers.LoadConfiguration("redis", "redis_pass")
	dbMaster := redis.NewClient(&redis.Options{
		Addr: host.(string),
		Password: pass.(string),
		DB: db,  // use default DB
	})

	return dbMaster
}
