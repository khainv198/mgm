package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
	db     *mongo.Database
}

func newClient(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

func New(ctx context.Context, dbName string, opts ...*options.ClientOptions) (*Client, error) {
	client, err := newClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &Client{client: client, db: db}, nil
}

func (c *Client) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	coll := c.db.Collection(name, opts...)

	return &Collection{Collection: coll}
}
