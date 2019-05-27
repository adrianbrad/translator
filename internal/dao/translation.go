package dao

import (
	"time"

	"golang.org/x/text/language"
)

type TranslationID int64

type Translation struct {
	ID        TranslationID `json:"translation_id"`
	CreatedAt time.Time     `json:"created_at"`
}

type TranslationText struct {
	TranslationID TranslationID `json:"translation_id"`
	LanguageTag   language.Tag  `json:"language_tag"`
	Text          string        `json:"text"`
}
