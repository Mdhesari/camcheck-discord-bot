package channelservice

import (
	"context"
	"mdhesari/camcheck-discord-bot/entity"
)

type CacheRepository interface {
	Create(key string, value string) (bool, error)
	Get(id string) (string, error)
	Delete(id string) (bool, error)
}

type Repository interface {
	Create(c context.Context, ch *entity.Channel) error
	GetAll(ctx context.Context, discordGID string) ([]entity.Channel, error)
	FindByDiscordID(ctx context.Context, id string) (*entity.Channel, error)
	RemoveChannelByDiscordID(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repo      Repository
	cacheRepo CacheRepository
}

func New(repo Repository, cacheRepo CacheRepository) Service {
	return Service{
		repo:      repo,
		cacheRepo: cacheRepo,
	}
}

func (s Service) Get(ctx context.Context, discordGID string) ([]entity.Channel, error) {
	channels, err := s.repo.GetAll(ctx, discordGID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
