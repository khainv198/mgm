package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDField struct {
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type DateFields struct {
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deletedAt"`
}

func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	return id, nil
}

func (f *IDField) GetID() interface{} {
	return f.ID
}

func (f *IDField) SetID(id interface{}) {
	f.ID = id.(*primitive.ObjectID)
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
