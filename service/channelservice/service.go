package channelservice

import (
	"context"
	"mdhesari/shawshank-discord-bot/entity"
)

type Repository interface {
	Create(c context.Context, ch *entity.Channel) error
	GetAll(ctx context.Context) ([]entity.Channel, error)
	FindByDiscordID(ctx context.Context, id string) (*entity.Channel, error)
	RemoveChannelByDiscordID(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Get(ctx context.Context) ([]entity.Channel, error) {
	channels, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
