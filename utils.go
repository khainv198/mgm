package mgm

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (c *Client) Coll(m Model, opts ...*options.CollectionOptions) *Collection {
	if collGetter, ok := m.(CollectionGetter); ok {
		return collGetter.Collection()
	}

	return c.Collection(CollName(m), opts...)
}

func CollName(m Model) string {
	if collNameGetter, ok := m.(CollectionNameGetter); ok {
		return collNameGetter.CollectionName()
	}

	name := reflect.TypeOf(m).Elem().Name()

	return inflection.Plural(toSnakeCase(name))
}

func UpsertTrueOption() *options.UpdateOptions {
	upsert := true
	return &options.UpdateOptions{Upsert: &upsert}
}

func ParseObjectId(inp interface{}) (*primitive.ObjectID, error) {
	if id, ok := inp.(primitive.ObjectID); ok {
		return &id, nil
	}

	if str, ok := inp.(string); ok {
		if id, err := primitive.ObjectIDFromHex(str); err == nil {
			return &id, nil
		}
	}

	id := primitive.NilObjectID

	return &id, errors.New("InvalidObjectID")
}

func NewObjectID() *primitive.ObjectID {
	id := primitive.NewObjectID()
	return &id
}
