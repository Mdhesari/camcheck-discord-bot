package messagehandler

import (
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	session *discordgo.Session
}

func New(s *discordgo.Session) *Handler {
	return &Handler{
		session: s,
	}
}

func (h Handler) SetHanlders() {
	h.session.AddHandler(h.ReplyCommands)
}