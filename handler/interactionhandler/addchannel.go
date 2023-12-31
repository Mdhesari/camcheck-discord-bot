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
	content := ""

	opt, ok := optionMap["channel"]
	if !ok {
		content = "Channel is invalid!"
		SendInteractionRespond(content, s, i)

		return
	}

	c := opt.ChannelValue(nil)

	if h.channelSrv.IsVideoChannel(context.Background(), c.ID) {
		content = "Channel already exists."

	} else {
		camcheckCh := entity.Channel{
			DiscordID: c.ID,
			GuildID:   i.GuildID,
			IsVideo:   true,
		}

		err := h.channelSrv.AddChannel(context.Background(), &camcheckCh)
		if err != nil {
			log.Println("Channel service add error: ", err)

			content = "> Something went wrong!"
		} else {
			content = "> channel added successfully: <#%s>\n"
		}

		content = fmt.Sprintf(content, c.ID)

	}

	SendInteractionRespond(content, s, i)
}
