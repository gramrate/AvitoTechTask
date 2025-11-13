package user

import (
	"AvitoTechTask/pkg/ent"
	"context"

	"github.com/google/uuid"
)

func (r *UserRepository) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}
