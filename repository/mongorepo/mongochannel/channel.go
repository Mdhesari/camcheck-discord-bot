package mongochannel

import (
	"context"
	"errors"
	"mdhesari/camcheck-discord-bot/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DB) GetAll(ctx context.Context, discordGID string) ([]entity.Channel, error) {
	ctx, cancel := context.WithTimeout(ctx, d.cli.QueryTimeout)
	defer cancel()

	var channels []entity.Channel
	cur, err := d.cli.Conn().Collection("channels").Find(ctx, bson.M{"guild_id": discordGID}, options.Find())
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var ch entity.Channel

		if err := cur.Decode(&ch); err != nil {
			panic(err)
		}

		channels = append(channels, ch)
	}

	return channels, nil
}

func (d *DB) FindByID(ctx context.Context, id string) (*entity.Channel, error) {
	ctx, cancel := context.WithTimeout(ctx, d.cli.QueryTimeout)
	defer cancel()

	var ch entity.Channel
	filter := bson.M{"_id": id}
	res := d.cli.Conn().Collection("channels").FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	res.Decode(&ch)

	return &ch, nil
}

func (d *DB) Create(ctx context.Context, ch *entity.Channel) error {
	ctx, cancel := context.WithTimeout(ctx, d.cli.QueryTimeout)
	defer cancel()

	result, err := d.cli.Conn().Collection("channels").InsertOne(ctx, ch)
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("Could not create a new channel")
	}

	return nil
}

func (d *DB) FindByDiscordID(ctx context.Context, id string) (*entity.Channel, error) {
	ctx, cancel := context.WithTimeout(ctx, d.cli.QueryTimeout)
	defer cancel()

	var ch entity.Channel
	filter := bson.M{"discord_id": id}
	res := d.cli.Conn().Collection("channels").FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	res.Decode(&ch)

	return &ch, nil
}

func (d *DB) RemoveChannelByDiscordID(ctx context.Context, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, d.cli.QueryTimeout)
	defer cancel()

	filter := bson.M{"discord_id": id}
	res, err := d.cli.Conn().Collection("channels").DeleteOne(ctx, filter)
	if err != nil {

		return false, err
	}

	if res.DeletedCount > 0 {

		return true, nil
	}

	return false, nil
}
