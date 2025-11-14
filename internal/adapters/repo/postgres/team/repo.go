package team

import (
	"AvitoTechTask/pkg/ent"
)

//type Repo interface {
//	Create(ctx context.Context, team *ent.Team) (*ent.Team, error)
//	Get(ctx context.Context, teamID uuid.UUID) (*ent.Team, error)
//	Update(ctx context.Context, team *ent.Team) (*ent.Team, error)
//	GetWithMembers(ctx context.Context, teamID uuid.UUID) (*ent.Team, error)
//}

type Repo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) *Repo {
	return &Repo{client: client}
}
