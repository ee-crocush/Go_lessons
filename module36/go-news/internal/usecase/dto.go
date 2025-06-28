// Package usecase выполняет бизнес-логику приложения.
package usecase

// ParsedRSSDTO представляет данные, извлечённые из RSS.
type ParsedRSSDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Link    string `json:"link"`
	PubTime int64  `json:"pub_time"`
}

// FindByIDInputDTO представляет входной DTO для поиска поста по ID.
type FindByIDInputDTO struct {
	ID int32 `json:"id"`
}

// FindByIDOutputDTO представляет выходной DTO для поиска поста по ID.
type FindByIDOutputDTO struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Link    string `json:"link"`
	PubTime string `json:"pub_time"`
}
