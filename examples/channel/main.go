package main

import (
	"context"
	"fmt"
	"log"
	"mdhesari/shawshank-discord-bot/config"
	"mdhesari/shawshank-discord-bot/repository/mongorepo"
	"mdhesari/shawshank-discord-bot/repository/mongorepo/mongochannel"
	"mdhesari/shawshank-discord-bot/service/channelservice"
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
	channelSrv := channelservice.New(repo)

	if channelSrv.IsVideoChannel(context.Background(), "hsddfdi") {
		fmt.Println("this is a video channel")
	} else {
		fmt.Println("not a video channel")
	}
}
