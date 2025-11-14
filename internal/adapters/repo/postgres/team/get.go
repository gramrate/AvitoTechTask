package team

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/team"
	"context"

	"github.com/google/uuid"
)

func (r *TeamRepository) Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	return r.client.Team.Get(ctx, teamID)
}

func (r *TeamRepository) GetWithMembers(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	return r.client.Team.Query().
		Where(team.IDEQ(teamID)).
		WithMembers().
		Only(ctx)
}
