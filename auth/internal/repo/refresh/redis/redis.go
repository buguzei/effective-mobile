package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type Redis struct {
	cli *redis.Client
}

func New() *Redis {
	cli := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})

	return &Redis{
		cli: cli,
	}
}

func (r Redis) SetRefresh(ctx context.Context, refreshToken string, id int) {
	r.cli.Set(ctx, strconv.Itoa(id), refreshToken, 0)
}
