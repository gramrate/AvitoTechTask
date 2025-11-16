package pull_request

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func (s *Service) ReassignReviewer(ctx context.Context, req *dto.ReassignPRRequest) (*dto.ReassignPRResponse, error) {
	// Получаем PR с ревьюверами
	pr, err := s.pullRequestsRepo.GetWithReviewers(ctx, req.PullRequestID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pull request: %w", err)
	}

	// Проверяем, что PR не в статусе MERGED
	if pr.Status == types.PullRequestStatusMerged {
		return nil, fmt.Errorf("%w: cannot reassign reviewers for merged pull request", errorz.ErrPRMerged)
	}
	fmt.Println(pr.AssignedReviewers)
	// Проверяем, что старый ревьювер действительно назначен на этот PR
	isOldReviewerAssigned := containsUUID(pr.AssignedReviewers, req.OldReviewerID)

	if !isOldReviewerAssigned {
		return nil, fmt.Errorf("%w: reviewer %s is not assigned to this pull request", errorz.ErrNotAssigned, req.OldReviewerID)
	}

	// Получаем доступных ревьюверов из команды заменяемого ревьювера
	availableReviewers, err := s.pullRequestsRepo.GetAvailableReviewersFromReviewerTeam(ctx, req.OldReviewerID, req.PullRequestID)
	if err != nil {
		return nil, fmt.Errorf("failed to get available reviewers: %w", err)
	}

	// Выбираем случайного ревьювера из доступных
	var newReviewerID uuid.UUID
	if len(availableReviewers) > 0 {
		newReviewerID = availableReviewers[rand.Intn(len(availableReviewers))].ID
	} else {
		return nil, fmt.Errorf("%w: no available reviewers from reviewer's team", errorz.ErrNoCandidate)
	}

	// Переназначаем ревьювера
	updatedPR, err := s.pullRequestsRepo.ReassignReviewer(ctx, req.PullRequestID, req.OldReviewerID, newReviewerID)
	if err != nil {
		return nil, fmt.Errorf("failed to reassign reviewer: %w", err)
	}

	// Конвертируем в DTO
	prDTO := s.convertToDTO(updatedPR)

	return &dto.ReassignPRResponse{
		PR: prDTO,
	}, nil
}

// convertToDTO конвертирует ent.PullRequest в dto.PullRequest
func (s *Service) convertToDTO(pr *ent.PullRequest) dto.PullRequest {
	dtoPR := dto.PullRequest{
		PullRequestID:   pr.ID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          pr.Status.String(),
	}

	// Добавляем назначенных ревьюверов
	if pr.Edges.Reviewers != nil {
		dtoPR.AssignedReviewers = make([]uuid.UUID, len(pr.Edges.Reviewers))
		for i, reviewer := range pr.Edges.Reviewers {
			dtoPR.AssignedReviewers[i] = reviewer.ID
		}
	}

	return dtoPR
}
func containsUUID(slice []uuid.UUID, item uuid.UUID) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
