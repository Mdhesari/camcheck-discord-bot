package main

import (
	"context"
	"fmt"
	"log"
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/repository/mongorepo"
	"mdhesari/camcheck-discord-bot/repository/mongorepo/mongochannel"
	"mdhesari/camcheck-discord-bot/repository/redisrepo"
	"mdhesari/camcheck-discord-bot/repository/redisrepo/redischannel"
	"mdhesari/camcheck-discord-bot/service/channelservice"
	"time"

	"github.com/hellofresh/janus/pkg/plugin/basic/encrypt"
)

func main() {
	cfg := config.Load("config.yml")

	cli, err := mongorepo.New(cfg.Database.MongoDB, 30*time.Second, encrypt.Hash{})
	if err != nil {
		log.Fatalf("mongo connect error %s", err)
	}
	repo := mongochannel.New(cli)

	redisCli, err := redisrepo.New(cfg.Database.Redis)
	if err != nil {
		panic(err)
	}
	CacheRepo := redischannel.New(redisCli)

	channelSrv := channelservice.New(repo, CacheRepo)

	if channelSrv.IsVideoChannel(context.Background(), "hsddfdi") {
		fmt.Println("this is a video channel")
	} else {
		fmt.Println("not a video channel")
	}
}
