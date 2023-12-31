package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Channel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DiscordID string             `json:"discord_id,omitempty" bson:"discord_id,omitempty"`
	GuildID   string             `json:"guild_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	IsVideo   bool               `json:"is_video,omitempty"`
}
