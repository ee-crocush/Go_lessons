package post

import (
	dom "GoNews/internal/domain/post"
	"context"
	"fmt"
)

// FindLastUseCase интерфейс для поиска последней новости.
type FindLastUseCase interface {
	Execute(ctx context.Context) (PostDTO, error)
}

type findLastUseCase struct {
	repo dom.Repository
}

// NewFindLastUseCase создает новый экземпляр usecase для поиска последней новости.
func NewFindLastUseCase(repo dom.Repository) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

func (uc *findLastUseCase) Execute(ctx context.Context) (PostDTO, error) {
	post, err := uc.repo.FindLast(ctx)
	if err != nil {
		return PostDTO{}, fmt.Errorf("FindLastUseCase.Execute: %w", err)
	}

	return PostDTO{
		ID:      post.ID().Value(),
		Title:   post.Title().Value(),
		Content: post.Content().Value(),
		Link:    post.Link().Value(),
		PubTime: post.PubTime().String(),
	}, nil
}
