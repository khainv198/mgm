package mgm

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	if err := callToBeforeCreateHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx, model, opts...)
	if err != nil {
		return err
	}

	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(ctx, model)
}

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx, filter, opts...).Decode(model)
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	if err := callToBeforeUpdateHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx, bson.M{"_id": model.GetID()}, bson.M{"$set": model}, opts...)
	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(ctx, res, model)
}

func delete(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	if err := callBeforeDeleteHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx, bson.M{"_id": model.GetID()}, bson.M{"$set": model}, opts...)
	if err != nil {
		return err
	}

	return callAfterDeleteHooks(ctx, res, model)
}

func restore(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	if err := callBeforeRestoreHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx, bson.M{"_id": model.GetID()}, bson.M{"$set": model}, opts...)
	if err != nil {
		return err
	}

	return callAfterRestoreHooks(ctx, res, model)
}

func destroy(ctx context.Context, c *Collection, model Model) error {
	if err := callBeforeDestroyHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.DeleteOne(ctx, bson.M{"_id": model.GetID()})
	if err != nil {
		return err
	}

	return callAfterDestroyHooks(ctx, res, model)
}

func get(ctx context.Context, c *Collection, pipeline []bson.M, result interface{}) error {
	cursor, err := c.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(result); err != nil {
			return err
		}

		return nil
	}

	return mongo.ErrNoDocuments
}

func list(ctx context.Context, c *Collection, pipeline []bson.M, results interface{}) error {
	cursor, err := c.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, results); err != nil {
		return err
	}

	return nil
}

func listAndCount(ctx context.Context, c *Collection, pipeline []bson.M, items interface{}, total *int64) error {
	var wg sync.WaitGroup
	wg.Add(2)

	var countErr, listErr error

	go func() {
		defer wg.Done()

		filter := bson.M{}
		if len(pipeline) > 0 {
			if pipeline[0]["$match"] != nil {
				if f, ok := pipeline[0]["$match"].(bson.M); ok {
					filter = f
				}
			}
		}

		*total, countErr = c.CountDocuments(ctx, filter)
	}()

	go func() {
		defer wg.Done()

		err := c.List(ctx, pipeline, items)
		if err != nil {
			listErr = err
		}
	}()

	wg.Wait()

	if countErr != nil {
		return countErr
	}

	if listErr != nil {
		return listErr
	}

	return nil
}
