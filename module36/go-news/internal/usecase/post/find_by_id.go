package post

import (
	"context"
	"fmt"

	dom "GoNews/internal/domain/post"
)

// FindByIDUseCase интерфейс для поиска новости по ID.
type FindByIDUseCase interface {
	Execute(ctx context.Context, in FindByIDInputDTO) (PostDTO, error)
}

type findByIDUseCase struct {
	repo dom.Repository
}

// NewFindByIDUseCase создает новый экземпляр usecase для поиска новости по ID.
func NewFindByIDUseCase(repo dom.Repository) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

// Execute выполняет бизнес-логику поиска новости по ID.
func (uc *findByIDUseCase) Execute(ctx context.Context, in FindByIDInputDTO) (PostDTO, error) {
	postID, err := dom.NewPostID(in.ID)
	if err != nil {
		return PostDTO{}, fmt.Errorf("FindByIDUseCase.NewPostID: %w", err)
	}

	post, err := uc.repo.FindByID(ctx, postID)
	if err != nil {
		return PostDTO{}, fmt.Errorf("findByIDUseCase.FindByID: %w", err)
	}

	return PostDTO{
		ID:      post.ID().Value(),
		Title:   post.Title().Value(),
		Content: post.Content().Value(),
		Link:    post.Link().Value(),
		PubTime: post.PubTime().String(),
	}, nil
}
