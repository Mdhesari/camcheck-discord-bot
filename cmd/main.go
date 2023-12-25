package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
	users []string
)

func init() {
	flag.StringVar(&token, "t", "default", "Discord token")
	flag.Parse()
}

func main() {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Discord session error: ", err)

		return
	}
	defer session.Close()

	session.AddHandler(func(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
		fmt.Println("user dancing..")
		if e.Member.User.Bot {

			return
		}

		if e.Member.User.ID == s.State.User.ID {

			return
		}

		if !e.SelfVideo && e.ChannelID != "" {
			// s.GuildMemberMove(e.GuildID, e.UserID, nil)
			users = append(users, e.UserID)
			go func(uid string) {
				time.Sleep(10 * time.Second)

				if len(users) > 0 {
					s.ChannelMessageSend(e.ChannelID, e.Member.User.Mention()+" Your camera is off! You will be disconnected very soon! that's all I know...")
				}

				time.Sleep(60 * time.Second)

				if err := s.GuildMemberMove(e.GuildID, uid, nil); err != nil {
					fmt.Println(err)
				}

				// TODO: remove only specified user
				users = []string{}
				// fmt.Println("video enabled: ", e.SelfVideo)
			}(e.UserID)
		} else if e.SelfVideo {
			users = []string{}
		}
	})

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "Ping" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong")
			if err != nil {
				fmt.Println("Message reply error: ", err)
			}
		}
	})

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
}
