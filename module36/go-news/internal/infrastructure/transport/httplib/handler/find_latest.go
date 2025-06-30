package handler

import (
	uc "GoNews/internal/usecase/post"
	"GoNews/pkg/api"
	"github.com/gofiber/fiber/v2"
)

// FindLatestRequest - входные данные из тела запроса для получения последних n новостей.
type FindLatestRequest struct {
	Limit int `json:"limit"`
}

// FindLatestResponse представляет ответ на запрос получения последних n новостей.
type FindLatestResponse struct {
	Posts []PostItem `json:"posts"`
}

// FindLatestHandler обрабатывает запрос (GET /news/latest).
func (h *Handler) FindLatestHandler(c *fiber.Ctx) error {
	req, err := api.Req[FindLatestRequest](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Err(err))
	}

	in := uc.FindLatestInputDTO{Limit: req.Limit}
	out, err := h.findLatestUC.Execute(c.Context(), in)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Err(err))
	}

	posts := make([]PostItem, len(out))
	for _, post := range out {
		posts = append(
			posts, PostItem{
				ID:      post.ID,
				Title:   post.Title,
				Content: post.Content,
				Link:    post.Link,
				PubTime: post.PubTime,
			},
		)
	}

	resp := FindAllResponse{
		Posts: posts,
	}

	return c.Status(fiber.StatusOK).JSON(api.Resp(resp))
}
