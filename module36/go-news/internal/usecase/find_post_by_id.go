package usecase

import (
	"context"
	"fmt"

	dom "GoNews/internal/domain/post"
)

// FindByIDUseCase интерфейс для поиска новости по ID.
type FindByIDUseCase interface {
	Execute(ctx context.Context, in FindByIDInputDTO) (FindByIDOutputDTO, error)
}

type findByIDUseCase struct {
	repo dom.Repository
}

// NewFindByIDUseCase создает новый экземпляр usecase для поиска новости по ID.
func NewFindByIDUseCase(repo dom.Repository) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

// Execute выполняет бизнес-логику поиска новости по ID.
func (f *findByIDUseCase) Execute(ctx context.Context, in FindByIDInputDTO) (FindByIDOutputDTO, error) {
	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return FindByIDOutputDTO{}, fmt.Errorf("FindByIDUseCase.NewPostID: %w", err)
	}

	post, err := f.repo.FindByID(ctx, postID)
	if err != nil {
		return FindByIDOutputDTO{}, fmt.Errorf("findByIDUseCase.FindByID: %w", err)
	}

	return FindByIDOutputDTO{
		ID:      post.ID().Value(),
		Title:   post.Title().Value(),
		Content: post.Content().Value(),
		Link:    post.Link().Value(),
		PubTime: post.PubTime().String(),
	}, nil
}
