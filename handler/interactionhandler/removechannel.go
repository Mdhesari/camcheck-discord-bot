package interactionhandler

import (
	"context"
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

	content := ""

	opt, ok := optionMap["channel"]

	if !ok {
		content = "Channel is invalid!"
		sendInteractionRespond(content, s, i)

		return
	}

	c := opt.ChannelValue(nil)

	isVideoChannel := h.channelSrv.IsVideoChannel(context.Background(), c.ID)

	if isVideoChannel {
		res, err := h.channelSrv.RemoveChannel(context.Background(), c.ID)
		if err != nil || !res {
			log.Println(err)

			content = "Something went wrong!"
		} else {

			content += "> channel deleted successfully: <#%s>\n"
		}
	} else {
		content = "> channel is not in camcheck's list: <#%s>\n"
	}

	sendInteractionRespond(content, s, i)
}

func sendInteractionRespond(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
