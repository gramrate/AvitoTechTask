package team

import (
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
)

func (r *Repo) Create(ctx context.Context, teamEntity *ent.Team) (*ent.Team, error) {
	team, err := r.client.Team.Create().
		SetTeamName(teamEntity.TeamName).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) || ent.IsValidationError(err) {
			// Проверяем, что ошибка связана с уникальностью имени команды
			return nil, fmt.Errorf("create team: %w", errorz.ErrTeamNameAlreadyUsed)
		}
		return nil, fmt.Errorf("create team: %w", err)
	}
	return team, nil
}
