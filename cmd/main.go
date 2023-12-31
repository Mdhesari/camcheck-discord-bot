package main

import (
	"flag"
	"log"
	"mdhesari/camcheck-discord-bot/config"
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

	"github.com/bwmarrin/discordgo"
	"github.com/hellofresh/janus/pkg/plugin/basic/encrypt"
)

var (
	cfg      config.Config
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "list-channel",
			Description: "Display bot channels list.",
		},
		{
			Name:        "add-channel",
			Description: "Adds a new channel to bot.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionChannel,
					Name: "channel",
					ChannelTypes: []discordgo.ChannelType{
						discordgo.ChannelTypeGuildVoice,
					},
					Description: "which channel to add.",
					Required:    true,
				},
			},
		},
		{
			Name:        "remove-channel",
			Description: "Removes a channel from bot.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionChannel,
					Name: "channel",
					ChannelTypes: []discordgo.ChannelType{
						discordgo.ChannelTypeGuildVoice,
					},
					Description: "which channel to remove.",
					Required:    true,
				},
			},
		},
	}
)

func init() {
	cfg = config.Load("config.yml")
	flag.Parse()
}

func main() {
	session, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		log.Fatalln("Discord session error: ", err)

		return
	}

	err = session.Open()
	if err != nil {
		log.Println("Open connection error: ", err)

		return
	}
	defer session.Close()

	session.Identify.Intents = discordgo.IntentsAll

	registerHandlers(session)

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	removeCommands(session, registeredCommands)
}

func removeCommands(s *discordgo.Session, registercmds []*discordgo.ApplicationCommand) {
	log.Println("Removing commands...")
	// // We need to fetch the commands, since deleting requires the command ID.
	// // We are doing this from the returned commands on line 375, because using
	// // this will delete all the commands, which might not be desirable, so we
	// // are deleting only the commands that we added.
	// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
	// if err != nil {
	// 	log.Fatalf("Could not fetch registered commands: %v", err)
	// }

	for _, v := range registercmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func registerHandlers(s *discordgo.Session) {
	cli, err := mongorepo.New(cfg.Database.MongoDB, 5*time.Second, encrypt.Hash{})
	if err != nil {
		panic(err)
	}
	repo := mongochannel.New(cli)

	redisCli, err := redisrepo.New(cfg.Database.Redis)
	if err != nil {
		panic(err)
	}
	CacheRepo := redischannelrepo.New(redisCli)

	channelSrv := channelservice.New(repo, CacheRepo)

	video := videohandler.New(&cfg.Discord, s, channelSrv)

	video.SetHandlers()

	message := messagehandler.New(s)

	message.SetHanlders()

	interaction := interactionhandler.New(&cfg.Discord, s, channelSrv)

	interaction.SetHandlers()
}
