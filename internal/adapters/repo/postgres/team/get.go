package team

import (
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/team"
	"context"

	"github.com/google/uuid"
)

func (r *Repo) Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	tm, err := r.client.Team.Get(ctx, teamID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrTeamNotFound
		}
		return nil, err
	}
	return tm, nil
}

func (r *Repo) GetByNameWithMembers(ctx context.Context, name string) (*ent.Team, error) {
	tm, err := r.client.Team.Query().
		Where(team.TeamNameEQ(name)).
		WithMembers().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrTeamNotFound
		}
		return nil, err
	}
	return tm, nil
}
