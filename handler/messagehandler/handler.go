package messagehandler

import (
	"mdhesari/camcheck-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config   *config.Discord
	handlers []func()
}

func New(cfg *config.Discord) *Handler {
	return &Handler{
		config: cfg,
	}
}

func (h Handler) Register(session *discordgo.Session) {
	actions := []interface{}{
		h.ReplyCommands,
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
