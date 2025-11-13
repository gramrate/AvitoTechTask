package team

import (
	"AvitoTechTask/pkg/ent"
	"context"

	"github.com/google/uuid"
)

func (r *TeamRepository) Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	return r.client.Team.Get(ctx, teamID)
}
