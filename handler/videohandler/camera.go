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

	if e.BeforeUpdate != nil {
		if e.BeforeUpdate.ChannelID != e.ChannelID && h.channelSrv.CamCheckUserIsOff(e.UserID, e.BeforeUpdate.ChannelID) {
			h.channelSrv.RemoveUserCamCheck(e.UserID)
		}
	}

	// check wether this video channel is watched
	if !h.channelSrv.IsVideoChannel(context.Background(), e.ChannelID) {

		return
	}

	log.Println("channel is in...")

	if !e.SelfVideo && e.ChannelID != "" {
		h.channelSrv.AddCamCheck(e.UserID, e.ChannelID)

		go func() {
			// give some time to enable camera the notice
			time.Sleep(10 * time.Second)

			if !h.channelSrv.CamCheckUserIsOff(e.UserID, e.ChannelID) {

				return
			}

			msg := fmt.Sprintf("%s Your camera is off! You will be disconnected very soon!\nThat's all I know...", e.Member.User.Mention())
			if _, err := s.ChannelMessageSend(e.ChannelID, msg); err != nil {
				log.Println("Failed to send channel message: ", err)
			}

			// last opportunity for user to enable camera
			time.Sleep(time.Duration(h.config.MaxWaitSeconds) * time.Second)

			if !h.channelSrv.CamCheckUserIsOff(e.UserID, e.ChannelID) {

				return
			}

			if err := s.GuildMemberMove(e.GuildID, e.UserID, nil); err != nil {
				log.Println("Failed to disconnect from channel: ", err)
			} else {
				h.channelSrv.RemoveUserCamCheck(e.UserID)
			}
		}()
	} else if e.SelfVideo {
		if h.channelSrv.CamCheckUserIsOff(e.UserID, e.ChannelID) {
			h.channelSrv.RemoveUserCamCheck(e.UserID)
		}
	}
}
