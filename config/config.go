package config

import "mdhesari/shawshank-discord-bot/repository/mongorepo"

type Database struct {
	MongoDB mongorepo.Config `koanf:"mongodb"`
}

type Discord struct {
	Token string
}

type Config struct {
	Database Database `koanf:"database"`
	Discord Discord `koanf:"discord"`
}
