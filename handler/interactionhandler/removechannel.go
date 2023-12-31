package interactionhandler

import (
	"context"
	"fmt"

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
		SendInteractionRespond(content, s, i)

		return
	}

	c := opt.ChannelValue(nil)

	isVideoChannel := h.channelSrv.IsVideoChannel(context.Background(), c.ID)

	if isVideoChannel {
		if res := h.channelSrv.RemoveChannelByDiscordID(context.Background(), c.ID); !res {
			content = "Something went wrong! <#%s>\n"
		} else {
			content = "> channel deleted successfully: <#%s>\n"
		}
	} else {
		content = "> channel is not in camcheck's list: <#%s>\n"
	}

	content = fmt.Sprintf(content, c.ID)

	SendInteractionRespond(content, s, i)
}