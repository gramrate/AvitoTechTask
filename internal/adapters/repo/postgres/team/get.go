package team

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/team"
	"context"

	"github.com/google/uuid"
)

func (r *Repo) Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	return r.client.Team.Get(ctx, teamID)
}

func (r *Repo) GetByNameWithMembers(ctx context.Context, name string) (*ent.Team, error) {
	return r.client.Team.Query().
		Where(team.TeamNameEQ(name)).
		WithMembers().
		Only(ctx)
}
