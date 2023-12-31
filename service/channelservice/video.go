package channelservice

import (
	"context"
	"log"
	"mdhesari/camcheck-discord-bot/entity"

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

func (s Service) RemoveChannelByDiscordID(ctx context.Context, id string) bool {
	res, err := s.repo.RemoveChannelByDiscordID(ctx, id)
	if err != nil {
		log.Println(err)

		return false
	}

	return res
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
		log.Println("Add user camera off error: ", err)

		return false
	}

	return true
}

func (s Service) RemoveUserCameraOff(channelID string, userID string) bool {
	err := s.cacheRepo.RemoveUserCameraOff(channelID, userID)
	if err != nil {
		log.Println("Remove user camera off error: ", err)

		return false
	}

	return true
}

func (s Service) IsUserCameraOn(channelID string, userID string) bool {
	res, err := s.cacheRepo.IsUserCameraOn(channelID, userID)
	if err != nil {
		log.Println("Redis repo error: ", err)

		return false
	}

	return res
}
