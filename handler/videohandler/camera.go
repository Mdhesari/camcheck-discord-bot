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
		h.channelSrv.AddUserCameraOff(e.ChannelID, e.UserID)

		go func(channelId string, userId string) {
			log.Println("self video: ", e.SelfVideo)

			time.Sleep(10 * time.Second)

			if !h.channelSrv.IsUserCameraOn(channelId, userId) {
				msg := fmt.Sprintf("%s Your camera is off! You will be disconnected very soon!\nThat's all I know...", e.Member.User.Mention())

				s.ChannelMessageSend(channelId, msg)
			}

			time.Sleep(time.Duration(h.config.MaxWaitSeconds) * time.Second)

			log.Println("self video: ", e.SelfVideo)

			if !h.channelSrv.IsUserCameraOn(channelId, userId) && e.ChannelID == channelId {
				if err := s.GuildMemberMove(channelId, userId, nil); err != nil {
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
