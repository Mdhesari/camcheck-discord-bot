package videohandler

import (
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config     *config.Discord
	channelSrv *channelservice.Service
	handlers   []func()
}

func New(cfg *config.Discord, chSrv *channelservice.Service) *Handler {
	return &Handler{
		config:     cfg,
		channelSrv: chSrv,
	}
}

func (h Handler) Register(session *discordgo.Session) {
	actions := []interface{}{
		h.CheckCameraAndDisconnect,
	}

	for _, a := range actions {
		h.handlers = append(h.handlers, session.AddHandler(a))
	}
}

func (h Handler) DeRegister(session *discordgo.Session) {
	for _, remove := range h.handlers {
		remove()
	}
}
