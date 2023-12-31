package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Channel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DiscordID string             `json:"discord_id,omitempty" bson:"discord_id,omitempty"`
	GuildID   string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	IsVideo   bool               `json:"is_video,omitempty" bson:"is_video,omitempty"`
}
