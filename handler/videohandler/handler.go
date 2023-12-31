package videohandler

import (
	"mdhesari/camcheck-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
}

type Handler struct {
	session    *discordgo.Session
	channelSrv channelservice.Service
}

func New(s *discordgo.Session, chs channelservice.Service) *Handler {
	return &Handler{
		session:    s,
		channelSrv: chs,
	}
}

func (h Handler) SetHandlers() {
	h.session.AddHandler(h.CheckCameraAndDisconnect)
}
