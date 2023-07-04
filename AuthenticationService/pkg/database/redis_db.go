package database

import (
	"AuthenticationService/pkg/database/tables"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Conn   *redis.Client
	Access tables.DataAccess
}

func (rdb *RedisDB) Connect() {
	rdb.Conn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
