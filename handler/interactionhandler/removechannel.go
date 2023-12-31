package interactionhandler

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) RemoveChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "remove-channel" {

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

		isVideoChannel := h.channelSrv.IsVideoChannel(context.Background(), c.ID)

		if isVideoChannel {
			res, err := h.channelSrv.RemoveChannel(context.Background(), c.ID)
			if err != nil || !res {
				log.Println(err)

				msgformat = "Something went wrong!"
			} else {

				msgformat += "> channel deleted successfully from shawshank: <#%s>\n"
			}
		} else {
			msgformat = "> channel is not in shawshank's list: <#%s>\n"
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
