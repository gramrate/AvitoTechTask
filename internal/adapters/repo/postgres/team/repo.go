package team

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/team"
	"context"

	"github.com/google/uuid"
)

//type TeamRepository interface {
//	Create(ctx context.Context, team *ent.Team) (*ent.Team, error)
//	Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error)
//	Update(ctx context.Context, team *ent.Team) (*ent.Team, error)
//	GetWithMembers(ctx context.Context, teamID uuid.UUID) (*ent.Team, error)
//}

type TeamRepository struct {
	client *ent.Client
}

func NewTeamRepository(client *ent.Client) *TeamRepository {
	return &TeamRepository{client: client}
}

func (r *TeamRepository) GetWithMembers(ctx context.Context, teamID uuid.UUID) (*ent.Team, error) {
	return r.client.Team.Query().
		Where(team.IDEQ(teamID)).
		WithMembers().
		Only(ctx)
}
