package mongochannel

import "mdhesari/camcheck-discord-bot/repository/mongorepo"

type DB struct {
	cli *mongorepo.MongoDB
}

func New(cli *mongorepo.MongoDB) *DB {
	return &DB{
		cli: cli,
	}
}
