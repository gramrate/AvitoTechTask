package user

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/user"
	"context"
	"fmt"
)

func (r *Repo) UpdateActivity(ctx context.Context, userEntity *ent.User) (*ent.User, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	// Обновляем активность пользователя
	update := tx.User.UpdateOneID(userEntity.ID).
		SetIsActive(userEntity.IsActive)

	// Если указана команда, обновляем связь
	if userEntity.Edges.Team != nil {
		update = update.SetTeamID(userEntity.Edges.Team.ID)
	}

	_, err = update.Save(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("updating user activity: %w", err))
	}

	// Получаем обновленного пользователя с загруженной командой
	updatedUser, err := tx.User.Query().
		Where(user.IDEQ(userEntity.ID)).
		WithTeam().
		Only(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("loading updated user with team: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return updatedUser, nil
}
