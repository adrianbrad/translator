package dao

import (
	"sync"

	"golang.org/x/text/language"
)

type translationHolder map[language.Tag]string

type translations map[language.Tag]map[string]translationHolder

type MemoryDAO struct {
	sync.RWMutex

	//map is of reference type so we can store the same map ar different tag/text combinations
	translations translations
}

func NewMemoryDAO() *MemoryDAO {
	return &MemoryDAO{
		translations: make(map[language.Tag]map[string]translationHolder),
	}
}

func (m *MemoryDAO) StoreTranslation(languageFrom language.Tag, textFrom string, languageTo language.Tag, textTo string) (err error) {
	m.Lock()
	defer m.Unlock()

	obj, translationFound := m.translations[languageFrom]
	if !translationFound {
		obj = make(map[string]translationHolder)
		m.translations[languageFrom] = obj
	}

	tHolder, ok := obj[textFrom]
	if !ok {
		tHolder = make(map[language.Tag]string)
		obj[textFrom] = tHolder
	}

	tHolder[languageTo] = textTo
	if !translationFound {
		tHolder[languageFrom] = textFrom
	}

	m.translations[languageTo] = map[string]translationHolder{
		textTo: tHolder,
	}

	return
}

func (m *MemoryDAO) GetTranslation(languageFrom language.Tag, wordFrom string, languageTo language.Tag) (wordTo string, err error) {
	obj, ok := m.translations[languageFrom]
	if !ok {
		err = newNoEntryForTheProvidedLanguageError(languageFrom)
		return
	}

	tHolder, ok := obj[wordFrom]
	if !ok {
		err = newNoEntryForTheProvidedWordInProvidedLanguage(wordFrom, languageFrom)
		return
	}

	wordTo, ok = tHolder[languageTo]
	if !ok {
		err = newNoEntryForTheProvidedWordInRequestedLanguage(wordFrom, languageTo)
		return
	}

	return
}
