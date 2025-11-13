package team

import (
	"AvitoTechTask/pkg/ent"
	"context"
)

func (r *TeamRepository) Create(ctx context.Context, teamEntity *ent.Team) (*ent.Team, error) {
	return r.client.Team.Create().
		SetTeamName(teamEntity.TeamName).
		Save(ctx)
}
