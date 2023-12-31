package config

import (
	"mdhesari/shawshank-discord-bot/repository/mongorepo"
	"mdhesari/shawshank-discord-bot/repository/redisrepo"
)

type Database struct {
	MongoDB mongorepo.Config `koanf:"mongodb"`
	Redis   redisrepo.Config `koanf:"redis"`
}

type Discord struct {
	Token string `koanf:"token"`
}

type Config struct {
	Database Database `koanf:"database"`
	Discord  Discord  `koanf:"discord"`
}

func New(db Database, discord Discord) Config {
	return Config{
		Database: db,
		Discord:  discord,
	}
}
