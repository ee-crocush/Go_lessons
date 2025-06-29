package post

import (
	dom "GoNews/internal/domain/post"
	"context"
	"fmt"
)

// FindLatestUseCase интерфейс для поиска последних n новостей.
type FindLatestUseCase interface {
	Execute(ctx context.Context, in FindLatestInputDTO) ([]PostDTO, error)
}

type findLatestUseCase struct {
	repo dom.Repository
}

// NewFindLatestUseCase создает новый экземпляр usecase для поиска последних n новостей.
func NewFindLatestUseCase(repo dom.Repository) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

// Execute выполняет бизнес-логику поиска последних n новостей.
func (uc *findLatestUseCase) Execute(ctx context.Context, in FindLatestInputDTO) ([]PostDTO, error) {
	in.Validate()

	posts, err := uc.repo.FindLatest(ctx, in.Limit)
	if err != nil {
		return []PostDTO{}, fmt.Errorf("FindLatestUseCase.Execute: %w", err)
	}

	return MapPostsToDTO(posts), nil
}
