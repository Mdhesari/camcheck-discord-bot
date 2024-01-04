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
		// when channel video is changed to voice channel we should consider that cause no longer need to notice!
		if !h.channelSrv.IsUserCameraOn(e.BeforeUpdate.ChannelID, e.UserID) {
			h.channelSrv.RemoveUserCameraOff(e.BeforeUpdate.ChannelID, e.UserID)
		}

		return
	}

	log.Println("channel is in...")

	if !e.SelfVideo && e.ChannelID != "" {
		h.channelSrv.AddUserCameraOff(e.ChannelID, e.UserID)

		go func() {
			// give some time to enable camera the notice
			time.Sleep(10 * time.Second)

			if !h.channelSrv.IsUserCameraOn(e.ChannelID, e.UserID) {
				msg := fmt.Sprintf("%s Your camera is off! You will be disconnected very soon!\nThat's all I know...", e.Member.User.Mention())

				if _, err := s.ChannelMessageSend(e.ChannelID, msg); err != nil {
					log.Println("Failed to send channel message: ", err)
				}
			}

			// last opportunity for user to enable camera
			time.Sleep(time.Duration(h.config.MaxWaitSeconds) * time.Second)

			if !h.channelSrv.IsUserCameraOn(e.ChannelID, e.UserID) {
				if err := s.GuildMemberMove(e.GuildID, e.UserID, nil); err != nil {
					log.Println("Failed to disconnect from channel: ", err)
				} else {
					h.channelSrv.RemoveUserCameraOff(e.ChannelID, e.UserID)
				}
			}
		}()
	} else if e.SelfVideo {
		if !h.channelSrv.IsUserCameraOn(e.ChannelID, e.UserID) {
			h.channelSrv.RemoveUserCameraOff(e.ChannelID, e.UserID)
		}
	}
}
