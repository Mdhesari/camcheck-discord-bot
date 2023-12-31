package main

import (
	"fmt"
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/repository/redisrepo"
	"mdhesari/camcheck-discord-bot/repository/redisrepo/redischannel"
)

func main() {
	cfg := config.Load("config.yml")
	redisCli, err := redisrepo.New(cfg.Database.Redis)
	if err != nil {
		panic(err)
	}
	CacheRepo := redischannel.New(redisCli)

	fmt.Println(CacheRepo.AddUserCameraOff("mychannel", "myuser"))

	fmt.Println(CacheRepo.IsUserCameraOn("mychannel", "myuser"))
}
