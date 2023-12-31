package interactionhandler

import (
	"mdhesari/shawshank-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	session    *discordgo.Session
	channelSrv *channelservice.Service
}

func New(s *discordgo.Session, channelSrv channelservice.Service) *Handler {
	return &Handler{
		session:    s,
		channelSrv: &channelSrv,
	}
}

func (h Handler) SetHandlers() {
	h.session.AddHandler(h.AddChannel)
	h.session.AddHandler(h.RemoveChannel)
	h.session.AddHandler(h.ListChannel)
}
