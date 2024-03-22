package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	*mongo.Collection
}

func (coll *Collection) FindByID(ctx context.Context, id interface{}, model Model, opts ...*options.FindOneOptions) error {
	id, err := model.PrepareID(id)
	if err != nil {
		return err
	}

	return first(ctx, coll, bson.M{"_id": id}, model, opts...)
}

func (coll *Collection) Create(ctx context.Context, model Model, opts ...*options.InsertOneOptions) error {
	return create(ctx, coll, model, opts...)
}

func (coll *Collection) Update(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return update(ctx, coll, model, opts...)
}

func (coll *Collection) Delete(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return delete(ctx, coll, model, opts...)
}

func (coll *Collection) Restore(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return restore(ctx, coll, model, opts...)
}

func (coll *Collection) Destroy(ctx context.Context, model Model) error {
	return destroy(ctx, coll, model)
}

func (coll *Collection) Get(ctx context.Context, pipeline []bson.M, result interface{}) error {
	return get(ctx, coll, pipeline, result)
}

func (coll *Collection) List(ctx context.Context, pipeline []bson.M, results interface{}) error {
	return list(ctx, coll, pipeline, results)
}

func (coll *Collection) ListAndCount(ctx context.Context, pipeline []bson.M, items interface{}, total *int64) error {
	return listAndCount(ctx, coll, pipeline, items, total)
}
