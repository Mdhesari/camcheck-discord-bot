package interactionhandler

import (
	"mdhesari/camcheck-discord-bot/config"
	"mdhesari/camcheck-discord-bot/service/channelservice"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config     *config.Discord
	channelSrv *channelservice.Service
	handlers   []func()
	commandIDs []string
}

func New(cfg *config.Discord, channelSrv *channelservice.Service) *Handler {
	return &Handler{
		config:     cfg,
		channelSrv: channelSrv,
	}
}

func (h Handler) Register(session *discordgo.Session) {
	h.TearUpCommands(session)

	actions := []interface{}{
		h.AddChannel,
		h.RemoveChannel,
		h.ListChannel,
		h.ManageChannels,
	}

	for _, a := range actions {
		h.handlers = append(h.handlers, session.AddHandler(a))
	}
}

func (h Handler) DeRegister(session *discordgo.Session) {
	h.TearDownCommands(session)

	for _, remove := range h.handlers {
		remove()
	}
}

func SendInteractionRespond(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
