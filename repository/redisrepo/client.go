package redisrepo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	DB       int    `koanf:"db"`
	Password string `koanf:"password"`
}

type Client struct {
	rdb    *redis.Client
	config Config
}

func New(cfg Config) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	status := cli.Ping(context.Background())
	if status.Err() != nil {

		return nil, status.Err()
	}

	return &Client{config: cfg, rdb: cli}, nil
}

func (c Client) RDB() *redis.Client {
	return c.rdb
}
