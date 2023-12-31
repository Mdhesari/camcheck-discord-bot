package channelservice

import (
	"context"
	"log"
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

func (s Service) RemoveChannel(ctx context.Context, id string) (bool, error) {
	return s.repo.RemoveChannelByDiscordID(ctx, id)
}

func (s Service) AddChannel(ctx context.Context, ch *entity.Channel) error {
	err := s.repo.Create(ctx, ch)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) AddUserCameraOff(channelID string, userID string) bool {
	err := s.cacheRepo.AddUserCameraOff(channelID, userID)
	if err != nil {
		log.Println(err)

		return false
	}

	return true
}

func (s Service) RemoveUserCameraOff(channelID string, userID string) bool {
	err := s.cacheRepo.RemoveUserCameraOff(channelID, userID)
	if err != nil {
		log.Println(err)

		return false
	}

	return true
}

func (s Service) IsUserCameraOn(channelID string, userID string) bool {
	res, err := s.cacheRepo.IsUserCameraOn(channelID, userID)
	if err != nil {
		log.Println("Reids repo error: ", err)

		return false
	}

	return res
}
