package videohandler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) CheckCameraAndDisconnect(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	if e.Member.User.ID == s.State.User.ID {

		return
	}

	// check wether this video channel is watched
	if !h.channelSrv.IsVideoChannel(context.Background(), e.ChannelID) {
		log.Println("not in...", e.ChannelID)

		return
	}

	log.Println("channel is in...")

	if !e.SelfVideo && e.ChannelID != "" {
		// s.GuildMemberMove(e.GuildID, e.UserID, nil)
		// users = append(users, e.UserID)
		go func(uid string) {
			time.Sleep(10 * time.Second)

			// check cache

			// check discord

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
}
