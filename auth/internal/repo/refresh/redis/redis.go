package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

type Redis struct {
	cli *redis.Client
}

func New() *Redis {
	cli := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	return &Redis{
		cli: cli,
	}
}

func (r Redis) GetRefresh(ctx context.Context, id int) (string, error) {
	res := r.cli.Get(ctx, strconv.Itoa(id))

	actualRefreshToken := res.Val()

	if err := res.Err(); err != nil {
		return "", err
	}

	return actualRefreshToken, nil
}

func (r Redis) SetRefresh(ctx context.Context, refreshToken string, id int) error {
	refreshTTL, err := time.ParseDuration(os.Getenv("REFRESH_TTL"))
	if err != nil {
		return err
	}

	r.cli.Set(ctx, strconv.Itoa(id), refreshToken, refreshTTL)
	return nil
}
