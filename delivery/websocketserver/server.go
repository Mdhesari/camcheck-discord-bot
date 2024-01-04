package websocketserver

import (
	"mdhesari/camcheck-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	Register(session *discordgo.Session)
	DeRegister(session *discordgo.Session)
}

type Server struct {
	Session  *discordgo.Session
	config   *config.Config
	handlers []Handler
}

func New(cfg *config.Config, handlers []Handler) *Server {
	s, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {

		panic(err)
	}

	// Todo: should be in config
	s.Identify.Intents = discordgo.IntentsAll

	return &Server{
		config:   cfg,
		Session:  s,
		handlers: handlers,
	}
}

func (s Server) Serve() {
	err := s.Session.Open()
	if err != nil {
		panic(err)
	}

	for _, h := range s.handlers {
		h.Register(s.Session)
	}
}

func (s Server) Shutdown() {
	for _, h := range s.handlers {
		h.DeRegister(s.Session)
	}

	s.Session.Close()
}
