package redischannel

import (
	"context"
	"mdhesari/shawshank-discord-bot/repository/redisrepo"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	cli *redisrepo.Client
}

func New(cli *redisrepo.Client) Repository {
	return Repository{
		cli: cli,
	}
}

func (r Repository) Create(id string) (bool, error) {
	res := r.cli.RDB().Set(context.Background(), id, id, time.Hour)

	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}

func (r Repository) Get(id string) (string, error) {
	res := r.cli.RDB().Get(context.Background(), id)

	if res.Err() != nil {
		return "", res.Err()
	}

	return res.Val(), nil
}

func (r Repository) Delete(id string) (bool, error) {
	res := r.cli.RDB().Unlink(context.Background(), id)
	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}

func (r Repository) IsUserCameraOn(channelID string, userID string) (bool, error) {
	res, err := r.Get(channelID + ":" + userID)
	if err != nil {
		if err == redis.Nil {
			return true, nil
		}

		return false, err
	}

	return res == "", nil
}

func (r Repository) AddUserCameraOff(channelID string, userID string) error {
	_, err := r.Create(channelID + ":" + userID)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) RemoveUserCameraOff(channelID string, userID string) error {
	_, err := r.Delete(channelID + ":" + userID)
	if err != nil {
		return err
	}

	return nil
}
