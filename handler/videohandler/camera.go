package videohandler

import (
	"context"
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
		h.channelSrv.AddUserCameraOff(e.ChannelID, e.UserID)

		go func(channelID string, userID string) {
			time.Sleep(30 * time.Second)

			if !h.channelSrv.IsUserCameraOn(channelID, userID) {
				s.ChannelMessageSend(e.ChannelID, e.Member.User.Mention()+" Your camera is off! You will be disconnected very soon! that's all I know...")
			}

			time.Sleep(10 * time.Second)

			if !h.channelSrv.IsUserCameraOn(channelID, userID) {
				if err := s.GuildMemberMove(e.GuildID, userID, nil); err != nil {
					log.Println("Failed to disconnect from channel: ", err)
				}
			}
		}(e.ChannelID, e.UserID)
	} else if e.SelfVideo {
		if !h.channelSrv.IsUserCameraOn(e.ChannelID, e.UserID) {
			h.channelSrv.RemoveUserCameraOff(e.ChannelID, e.UserID)
		}
	}
}
