package user

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/user"
	"context"
	"fmt"
)

func (r *Repo) CreateWithTeam(ctx context.Context, userEntity *ent.User) (*ent.User, error) {
	if userEntity.Edges.Team == nil {
		return nil, fmt.Errorf("userEntity.Edges.Team must be set")
	}
	teamID := userEntity.Edges.Team.ID

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	createdUser, err := tx.User.Create().
		SetID(userEntity.ID).
		SetUsername(userEntity.Username).
		SetIsActive(userEntity.IsActive).
		SetTeamID(teamID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("creating user: %w", err))
	}

	userWithTeam, err := tx.User.Query().
		Where(user.IDEQ(createdUser.ID)).
		WithTeam().
		Only(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("loading user with team: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return userWithTeam, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
	}
	return err
}
