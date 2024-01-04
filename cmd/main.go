package main

import (
	"flag"
	"log"
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/delivery/websocketserver"
	"mdhesari/camcheck-discord-bot/handler/interactionhandler"
	"mdhesari/camcheck-discord-bot/handler/messagehandler"
	"mdhesari/camcheck-discord-bot/handler/videohandler"
	"mdhesari/camcheck-discord-bot/repository/mongorepo"
	"mdhesari/camcheck-discord-bot/repository/mongorepo/mongochannel"
	"mdhesari/camcheck-discord-bot/repository/redisrepo"
	"mdhesari/camcheck-discord-bot/repository/redisrepo/redischannelrepo"
	"mdhesari/camcheck-discord-bot/service/channelservice"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hellofresh/janus/pkg/plugin/basic/encrypt"
)

var (
	cfg config.Config
)

func init() {
	cfg = config.Load("config.yml")
	flag.Parse()
}

func main() {
	server := websocketserver.New(&cfg, getHandlers())

	server.Serve()
	defer server.Shutdown()

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func getHandlers() []websocketserver.Handler {
	cli, err := mongorepo.New(cfg.Database.MongoDB, 5*time.Second, encrypt.Hash{})
	if err != nil {
		panic(err)
	}

	channeldbrepo := mongochannel.New(cli)

	redisCli, err := redisrepo.New(cfg.Database.Redis)
	if err != nil {
		panic(err)
	}

	channelcacherepo := redischannelrepo.New(redisCli)

	channelSrv := channelservice.New(channeldbrepo, channelcacherepo)

	return []websocketserver.Handler{
		videohandler.New(&cfg.Discord, &channelSrv),
		messagehandler.New(&cfg.Discord),
		interactionhandler.New(&cfg.Discord, &channelSrv),
	}
}
