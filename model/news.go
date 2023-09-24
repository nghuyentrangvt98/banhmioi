package model

import (
	"time"

	"github.com/nghuyentrangvt98/banhmioi/ent"
)

type News struct {
	ID         int       `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Subtitle   string    `json:"subtitle,omitempty"`
	Content    string    `json:"content,omitempty"`
	Author     string    `json:"author,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Category   string    `json:"category,omitempty"`
	ProductURL string    `json:"product_url,omitempty"`
	ImageURL   string    `json:"image_url,omitempty"`
}

type NewsList struct {
	Data  []*ent.News `json:"data,omitempty"`
	Total int         `json:"total,omitempty"`
}

type NewsQuery struct {
	Category string `query:"category"`
	Keyword  string `query:"keyword"`
}
