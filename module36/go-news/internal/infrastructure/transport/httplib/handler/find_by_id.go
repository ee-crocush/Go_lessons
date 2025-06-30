package handler

import (
	uc "GoNews/internal/usecase/post"
	"GoNews/pkg/api"
	"github.com/gofiber/fiber/v2"
)

// FindByIdRequest - входные данные из тела запроса для получения новости по ID.
type FindByIdRequest struct {
	ID int32 `json:"id"`
}

// FindByIdRequestResponse представляет выходной DTO поста.
type FindByIdRequestResponse struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Link    string `json:"link"`
	PubTime string `json:"pub_time"`
}

// FindByIDHandler обрабатывает запрос (GET /news/<id>).
func (h *Handler) FindByIDHandler(c *fiber.Ctx) error {
	req, err := api.Req[FindByIdRequest](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Err(err))
	}

	in := uc.FindByIDInputDTO{ID: req.ID}
	out, err := h.findByIDUC.Execute(c.Context(), in)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Err(err))
	}

	resp := FindByIdRequestResponse{
		ID:      out.ID,
		Title:   out.Title,
		Content: out.Content,
		Link:    out.Link,
		PubTime: out.PubTime,
	}

	return c.Status(fiber.StatusOK).JSON(api.Resp(resp))
}
