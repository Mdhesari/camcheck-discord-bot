package interactionhandler

import (
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config     *config.Discord
	session    *discordgo.Session
	channelSrv *channelservice.Service
}

func New(cfg *config.Discord, s *discordgo.Session, channelSrv channelservice.Service) *Handler {
	return &Handler{
		config:     cfg,
		session:    s,
		channelSrv: &channelSrv,
	}
}

func (h Handler) SetHandlers() {
	h.session.AddHandler(h.AddChannel)
	h.session.AddHandler(h.RemoveChannel)
	h.session.AddHandler(h.ListChannel)
}
