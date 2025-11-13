package user

import (
	"AvitoTechTask/pkg/ent"
	"context"
)

func (r *UserRepository) Create(ctx context.Context, userEntity *ent.User) (*ent.User, error) {
	return r.client.User.Create().
		SetID(userEntity.ID).
		SetUsername(userEntity.Username).
		SetIsActive(userEntity.IsActive).
		Save(ctx)
}
