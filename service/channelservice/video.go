package channelservice

import (
	"context"
	"mdhesari/shawshank-discord-bot/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s Service) IsVideoChannel(ctx context.Context, id string) bool {
	_, err := s.repo.FindByDiscordID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return false
		} else {

			panic(err)
		}
	}

	return true
}

func (s Service) AddChannel(ctx context.Context, ch *entity.Channel) error {
	err := s.repo.Create(ctx, ch)
	if err != nil {
		return err
	}

	return nil
}
