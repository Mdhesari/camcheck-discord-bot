package redischannelrepo

import (
	"context"
	"mdhesari/camcheck-discord-bot/repository/redisrepo"
	"time"
)

type Repository struct {
	cli *redisrepo.Client
}

func New(cli *redisrepo.Client) Repository {
	return Repository{
		cli: cli,
	}
}

func (r Repository) Create(key string, value string) (bool, error) {
	res := r.cli.RDB().Set(context.Background(), key, value, time.Hour)

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
