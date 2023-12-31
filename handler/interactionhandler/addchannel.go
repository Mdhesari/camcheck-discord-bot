package interactionhandler

import (
	"context"
	"fmt"
	"log"
	"mdhesari/camcheck-discord-bot/entity"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) AddChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "add-channel" {

		return
	}

	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	// format the bot's response
	margs := make([]interface{}, 0, len(options))
	msgformat := ""

	if opt, ok := optionMap["channel"]; ok {
		c := opt.ChannelValue(nil)

		margs = append(margs, opt.ChannelValue(nil).ID)

		camcheckCh := entity.Channel{
			DiscordID: c.ID,
			GuildID:   i.GuildID,
			IsVideo:   true,
		}
		err := h.channelSrv.AddChannel(context.Background(), &camcheckCh)
		if err != nil {
			log.Println("Channel service add error: ", err)

			msgformat = "> Something went wrong!"
		} else {
			msgformat += "> channel added successfully: <#%s>\n"
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				msgformat,
				margs...,
			),
		},
	})
}
