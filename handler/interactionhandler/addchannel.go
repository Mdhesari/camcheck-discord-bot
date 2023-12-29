package interactionhandler

import (
	"context"
	"fmt"
	"log"
	"mdhesari/shawshank-discord-bot/entity"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
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
	// msgformat := "You learned how to use command options! " +
	// 	"Take a look at the value(s) you entered:\n"

	// Get the value from the option map.
	// When the option exists, ok = true
	if opt, ok := optionMap["channel"]; ok {
		c := opt.ChannelValue(nil)

		margs = append(margs, opt.ChannelValue(nil).ID)

		shawshankch := entity.Channel{
			ID:        uuid.NewString(),
			DiscordID: c.ID,
			GuildID:   c.GuildID,
			Name:      c.Name,
			IsVideo:   true,
		}
		err := h.channelSrv.AddChannel(context.Background(), &shawshankch)
		if err != nil {
			log.Println(err)

			msgformat = "Something went wrong!"
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
