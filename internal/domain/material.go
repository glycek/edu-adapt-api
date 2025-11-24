package domain

import (
	"time"
)

// MaterialType — тип материала (чтобы не писать строки руками)
type MaterialType string

const (
	TypeBook    MaterialType = "book"
	TypeVideo   MaterialType = "video"
	TypeArticle MaterialType = "article"
)

// Material — главная структура
type Material struct {
	ID          int64        `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`             // Заголовок
	Description string       `json:"description" db:"description"` // Описание
	Type        MaterialType `json:"type" db:"type"`               // book/video/article

	// Метаданные для поиска
	DifficultyLevel string   `json:"difficulty_level" db:"difficulty_level"`
	Tags            []string `json:"tags" db:"tags"` // Теги (Postgres Array)

	FileURL   string    `json:"file_url" db:"file_url"` // Путь к файлу
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
