package user

import (
	"AvitoTechTask/pkg/ent"
)

//type UserRepository interface {
//	Create(ctx context.Context, user *ent.User) (*ent.User, error)
//	Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
//	Update(ctx context.Context, user *ent.User) (*ent.User, error)
//	GetByUsername(ctx context.Context, username string) (*ent.User, error)
//}

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{client: client}
}
