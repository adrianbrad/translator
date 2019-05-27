package dao

import (
	"golang.org/x/text/language"
)

type translationHolder map[language.Tag]string

type translations map[language.Tag]map[string]*translationHolder

type MemoryDAO struct {
	translations translations
}

func NewMemoryDAO() *MemoryDAO {
	return &MemoryDAO{
		translations: make(map[language.Tag]map[string]*translationHolder),
	}
}
