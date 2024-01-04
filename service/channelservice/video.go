package channelservice

import (
	"context"
	"log"
	"mdhesari/camcheck-discord-bot/entity"

	"github.com/redis/go-redis/v9"
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

func (s Service) AddCamCheck(userID string, channelID string) bool {
	res, err := s.cacheRepo.Create(userID, channelID)
	if err != nil {
		log.Println("Add user camera off error: ", err)

		return false
	}

	return res
}

func (s Service) RemoveUserCamCheck(userID string) bool {
	res, err := s.cacheRepo.Delete(userID)
	if err != nil {
		log.Println("Remove user camera off error: ", err)

		return false
	}

	return res
}

func (s Service) CamCheckUserIsOff(userID string, channelID string) bool {
	id, err := s.cacheRepo.Get(userID)
	if err != nil {
		if err == redis.Nil {
			return false
		}

		log.Println("Redis repo error: ", err)

		return false
	}

	return id == channelID
}

func (s Service) camCheckUser(userID string) bool {
	_, err := s.cacheRepo.Get(userID)
	if err != nil {
		if err == redis.Nil {
			return false
		}

		log.Println("Redis repo error: ", err)

		return false
	}

	return true
}
