package entity

type Channel struct {
	ID        string `json:"id"`
	DiscordID string `json:"discord_id"`
	GuildID   string `json:"guild_id"`
	Name      string `json:"name"`
	IsVideo   bool   `json:"is_video"`
}
