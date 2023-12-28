package main

import (
	"flag"
	"fmt"
	"log"
	"mdhesari/shawshank-discord-bot/handler"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token    = flag.String("token", "", "Discord token.")
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	users    []string
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "add-channel",
			Description: "Adds a new channel to shawshank.",
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
			Description: "Removes a channel from shawshank.",
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
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add-channel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// format the bot's response
			margs := make([]interface{}, 0, len(options))
			msgformat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"

			// Get the value from the option map.
			// When the option exists, ok = true
			if opt, ok := optionMap["channel"]; ok {
				margs = append(margs, opt.ChannelValue(nil).ID)
				msgformat += "> channel-option: <#%s>\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
	}
)

func init() {
	flag.Parse()
}

func main() {
	session, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatalln("Discord session error: ", err)

		return
	}
	defer session.Close()

	registeredcmds := registerCommands(session)

	registerHandlers(session)

	session.Identify.Intents = discordgo.IntentsAll

	err = session.Open()
	if err != nil {
		fmt.Println("Open connection error: ", err)

		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	removeCommands(session, registeredcmds)
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

func registerCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands
}

func registerHandlers(s *discordgo.Session) {
	video := handler.Video{}

	s.AddHandler(video.CheckCameraAndDisconnect)

	message := handler.Message{}

	s.AddHandler(message.ReplyCommands)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
