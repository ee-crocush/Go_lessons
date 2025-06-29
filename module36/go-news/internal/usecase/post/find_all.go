package post

import (
	dom "GoNews/internal/domain/post"
	"context"
	"fmt"
)

// FindAllUseCase интерфейс для поиска всех новостей.
type FindAllUseCase interface {
	Execute(ctx context.Context) ([]PostDTO, error)
}

type findAllUseCase struct {
	repo dom.Repository
}

// NewFindAllUseCase создает новый экземпляр usecase для поиска всех новостей.
func NewFindAllUseCase(repo dom.Repository) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

// Execute выполняет бизнес-логику поиска всех новостей.
func (uc *findAllUseCase) Execute(ctx context.Context) ([]PostDTO, error) {
	posts, err := uc.repo.FindAll(ctx)
	if err != nil {
		return []PostDTO{}, fmt.Errorf("FindAllUseCase.Execute: %w", err)
	}

	return MapPostsToDTO(posts), nil
}
