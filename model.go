package mgm

type CollectionGetter interface {
	Collection() *Collection
}

type CollectionNameGetter interface {
	CollectionName() string
}

type Model interface {
	PrepareID(id interface{}) (interface{}, error)

	GetID() interface{}
	SetID(id interface{})
}

type DefaultModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

func (model *DefaultModel) Creating() error {
	return model.DateFields.Creating()
}

func (model *DefaultModel) Saving() error {
	return model.DateFields.Saving()
}

func (model *DefaultModel) Deleting() error {
	return model.DateFields.Deleting()
}

func (model *DefaultModel) Restoring() error {
	return model.DateFields.Restoring()
}
