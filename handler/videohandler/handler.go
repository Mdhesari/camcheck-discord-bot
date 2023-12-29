package videohandler

import (
	"mdhesari/shawshank-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

var users []string

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
