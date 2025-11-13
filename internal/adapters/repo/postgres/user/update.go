package user

import (
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UserRepository) Update(ctx context.Context, userEntity *ent.User) (*ent.User, error) {
	if userEntity.ID == uuid.Nil {
		return nil, fmt.Errorf("user ID is required for update")
	}

	update := r.client.User.UpdateOneID(userEntity.ID).
		SetUsername(userEntity.Username).
		SetIsActive(userEntity.IsActive)

	if userEntity.Edges.Team != nil {
		update.SetTeamID(userEntity.Edges.Team.ID)
	}

	return update.Save(ctx)
}
