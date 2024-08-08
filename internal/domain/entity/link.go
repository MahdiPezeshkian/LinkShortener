package entity

import (
	"time"
)

type Link struct {
	Id          string    `json:"id"`
	Isdeleted   bool      `json:"is_deleted"`
	IsVisibled  bool      `json:"is_visibled"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Expiration  time.Time `json:"expiration"`
	Clicks      int       `json:"clicks"`
}
