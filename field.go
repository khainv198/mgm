package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDField struct {
	ID *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

type DateFields struct {
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return ParseObjectId(idStr)
	}

	return id, nil
}

func (f *IDField) GetID() interface{} {
	return f.ID
}

func (f *IDField) SetID(id interface{}) {
	ID := id.(primitive.ObjectID)
	f.ID = &ID
}

func (f *DateFields) Creating() error {
	now := time.Now().UTC()
	f.CreatedAt = &now
	return nil
}

func (f *DateFields) Saving() error {
	now := time.Now().UTC()
	f.UpdatedAt = &now
	return nil
}

func (f *DateFields) Deleting() error {
	now := time.Now().UTC()
	f.DeletedAt = &now
	return nil
}

func (f *DateFields) Restoring() error {
	f.DeletedAt = nil
	return nil
}
