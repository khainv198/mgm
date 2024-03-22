package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CreatingHook interface {
	Creating() error
}

type CreatingHookWithCtx interface {
	Creating(context.Context) error
}

type CreatedHook interface {
	Created() error
}

type CreatedHookWithCtx interface {
	Created(context.Context) error
}

type SavingHook interface {
	Saving() error
}

type SavingHookWithCtx interface {
	Saving(context.Context) error
}

type SavedHook interface {
	Saved() error
}

type SavedHookWithCtx interface {
	Saved(context.Context) error
}

func callToBeforeCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatingHookWithCtx); ok {
		if err := hook.Creating(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(CreatingHook); ok {
		if err := hook.Creating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatedHookWithCtx); ok {
		if err := hook.Created(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(CreatedHook); ok {
		if err := hook.Created(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHookWithCtx); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

// * Update
type UpdatingHook interface {
	Updating() error
}

type UpdatingHookWithCtx interface {
	Updating(context.Context) error
}

type UpdatedHook interface {
	Updated(result *mongo.UpdateResult) error
}

type UpdatedHookWithCtx interface {
	Updated(ctx context.Context, result *mongo.UpdateResult) error
}

func callToBeforeUpdateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(UpdatingHookWithCtx); ok {
		if err := hook.Updating(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(UpdatingHook); ok {
		if err := hook.Updating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterUpdateHooks(ctx context.Context, updateResult *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(UpdatedHookWithCtx); ok {
		if err := hook.Updated(ctx, updateResult); err != nil {
			return err
		}
	} else if hook, ok := model.(UpdatedHook); ok {
		if err := hook.Updated(updateResult); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHookWithCtx); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

// * Destroy
type DestroyingHook interface {
	Destroying() error
}

type DestroyingHookWithCtx interface {
	Destroying(context.Context) error
}

type DestroyedHook interface {
	Destroyed(*mongo.DeleteResult) error
}

type DestroyedHookWithCtx interface {
	Destroyed(context.Context, *mongo.DeleteResult) error
}

func callBeforeDestroyHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(DestroyingHookWithCtx); ok {
		if err := hook.Destroying(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(DestroyingHook); ok {
		if err := hook.Destroying(); err != nil {
			return err
		}
	}

	return nil
}

func callAfterDestroyHooks(ctx context.Context, deleteResult *mongo.DeleteResult, model Model) error {
	if hook, ok := model.(DestroyedHookWithCtx); ok {
		if err := hook.Destroyed(ctx, deleteResult); err != nil {
			return err
		}
	} else if hook, ok := model.(DestroyedHook); ok {
		if err := hook.Destroyed(deleteResult); err != nil {
			return err
		}
	}

	return nil
}

// * delete (soft delete)
type DeletingHook interface {
	Deleting() error
}

type DeletingHookWithCtx interface {
	Deleting(context.Context) error
}

type DeletedHook interface {
	Deleted(result *mongo.UpdateResult) error
}

type DeletedHookWithCtx interface {
	Deleted(ctx context.Context, result *mongo.UpdateResult) error
}

func callBeforeDeleteHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(DeletingHookWithCtx); ok {
		if err := hook.Deleting(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(DeletingHook); ok {
		if err := hook.Deleting(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callAfterDeleteHooks(ctx context.Context, update *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(DeletedHookWithCtx); ok {
		if err := hook.Deleted(ctx, update); err != nil {
			return err
		}
	} else if hook, ok := model.(DeletedHook); ok {
		if err := hook.Deleted(update); err != nil {
			return err
		}
	}

	return nil
}

// * restore
type RestoringHook interface {
	Restoring() error
}

type RestoringHookWithCtx interface {
	Restoring(context.Context) error
}

type RestoredHook interface {
	Restored(*mongo.UpdateResult) error
}

type RestoredHookWithCtx interface {
	Restored(context.Context, *mongo.UpdateResult) error
}

func callBeforeRestoreHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(RestoringHookWithCtx); ok {
		if err := hook.Restoring(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(RestoringHook); ok {
		if err := hook.Restoring(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callAfterRestoreHooks(ctx context.Context, result *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(RestoredHookWithCtx); ok {
		if err := hook.Restored(ctx, result); err != nil {
			return err
		}
	} else if hook, ok := model.(RestoredHook); ok {
		if err := hook.Restored(result); err != nil {
			return err
		}
	}

	return nil
}
