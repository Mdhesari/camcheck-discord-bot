package interactionhandler

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) ListChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "list-channel" {

		return
	}

	channels, err := h.channelSrv.Get(context.Background(), i.GuildID)
	if err != nil {
		log.Fatalf("Getting channels error: %s", err)

		return
	}

	content := fmt.Sprintf("Here's %s channels list :\n", h.config.Name)

	for _, ch := range channels {
		content += fmt.Sprintf("> channel : <#%s>\n", ch.DiscordID)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
