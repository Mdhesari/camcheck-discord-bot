package videohandler

import (
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config     *config.Discord
	session    *discordgo.Session
	channelSrv channelservice.Service
}

func New(cfg *config.Discord, s *discordgo.Session, chs channelservice.Service) *Handler {
	return &Handler{
		config:     cfg,
		session:    s,
		channelSrv: chs,
	}
}

func (h Handler) SetHandlers() {
	h.session.AddHandler(h.CheckCameraAndDisconnect)
}
