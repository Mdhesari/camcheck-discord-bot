package interactionhandler

import "github.com/bwmarrin/discordgo"

func (h Handler) ManageChannels(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "camcheck" {

		return
	}

	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Manage available channels to check: ",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:     discordgo.ChannelSelectMenu,
							CustomID:     "channel-selector",
							Placeholder:  "Pick your video channels.",
							MinValues:    new(int),
							MaxValues:    0,
							Options:      []discordgo.SelectMenuOption{},
							ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildText},
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}

	s.InteractionRespond(i.Interaction, &resp)
}
