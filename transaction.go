package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionFunc func(session mongo.Session, sc mongo.SessionContext) error

func (c *Client) Transaction(ctx context.Context, f TransactionFunc) error {
	return TransactionWithClient(ctx, c.client, f)
}

func TransactionWithClient(ctx context.Context, client *mongo.Client, f TransactionFunc) error {
	session, err := client.StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil { // startTransaction need to get options.
		return err
	}

	wrapperFn := func(sc mongo.SessionContext) error {
		return f(session, sc)
	}

	return mongo.WithSession(ctx, session, wrapperFn)
}
